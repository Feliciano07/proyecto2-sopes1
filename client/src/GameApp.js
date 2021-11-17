import React, { useEffect } from 'react'
import {BrowserRouter as Router, Switch, Route, Link} from 'react-router-dom'
import {HomePage} from './pages/HomePage'
import {TopJuegos} from './pages/TopJuegos'
import {TopJugadores} from './pages/TopJugadores'
import {UltimosJuegosPage} from './pages/UltimosJuegosPage';
import {WorkerPage} from './pages/WorkerPage'
import io from 'socket.io-client'
let socket;

const GameApp = ()=>{

    let direccion = process.env.REACT_APP_BACKEND_URL || "https://socket-rku52sdvga-wn.a.run.app"
    const ENDPOINT = direccion;
    
    console.log(process.env.REACT_APP_BACKEND_URL)
    console.log(direccion)

    socket = io(ENDPOINT,{
      reconnect: true,
      secure: true,
      transports: ['websocket', 'polling']
    });

    useEffect(() => {
      return () => socket.disconnect();
    }, []);

    return (

    <Router>
      <nav className="navbar navbar-dark bg-primary navbar-expand-sm fixed-top">
        <div className="container">
          <h3 className="navbar-brand" >Sistemas Operativos 1</h3>
          <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#Navbar">
                <span className="navbar-toggler-icon"></span>
            </button>

          <div className="collapse navbar-collapse" id="Navbar">
            <ul className="navbar-nav ml-auto">
              <li className="nav-item active">
                <Link to="/" className="nav-link" >Datos</Link>
              </li>
              <li className="nav-item">
                <Link to="/top3" className="nav-link">Top 3 Juegos</Link>
              </li>
              <li className="nav-item">
                <Link to="/worker" className="nav-link">Graficas Worker</Link>
              </li>
              <li className="nav-item">
                <Link to="/juegos" className="nav-link">Ultimos Juegos</Link>
              </li>
              <li className="nav-item">
                <Link to="/jugadores" className="nav-link">Top Mejores Jugadores</Link>
              </li>
            </ul>
          </div>
        </div>
      </nav>
      <br/><br/><br/>
      
      <div className="container pt-4">
        <Switch>
          <Route path="/top3">
            <TopJuegos socket = {socket}/>
          </Route>
          <Route path="/worker">
            <WorkerPage socket = {socket}/>
          </Route>
          <Route path="/juegos">
            <UltimosJuegosPage socket = {socket}/>
          </Route>
          <Route path="/jugadores">
            <TopJugadores socket = {socket}/>
          </Route>
          <Route path="/">
            <HomePage socket ={socket}/>
          </Route>
        </Switch>
      </div>
    </Router>

    )
}

export default GameApp;