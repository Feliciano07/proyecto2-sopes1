import React, { useState, useEffect } from 'react'

export const UltimosJuegosPage = ({socket}) => {

    const [datos, setDatos] = useState([]);

    useEffect(() => {
        socket.on("lastGame", data =>  setDatos(data) );
        console.log(datos);
      }, []);

    return (
        <div className="row pt-4" >
            <table className="table table-hover">
                    <thead>
                        <tr className="table-info">
                        <th scope="col">Id Game</th>
                        <th scope="col">Game Name</th>
                        <th scope="col">Winner</th>
                        <th scope="col">Worker</th>
                        </tr>
                    </thead>
                    <tbody>
                        {
                            datos.map( ( juego ,i) => {
                                return (
                                    <tr key={i}>
                                        <th scope="row">{ juego.game }</th>
                                        <th >{ juego.gamename }</th>
                                        <th >{ juego.winner }</th>
                                        <th >{ juego.worker }</th>
                                    </tr>
                                );
                            }) 
                        }
                        
                    </tbody>
                </table>
        </div>
    )
}
