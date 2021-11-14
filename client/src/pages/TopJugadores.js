import React, { useState, useEffect } from "react";

export const TopJugadores = ({socket}) => {

    const [datos, setDatos] = useState([]);
    
    useEffect(() => {
      socket.on("mejores", data =>  setDatos(data) );
    }, []);

    return (
        <div className="row pt-4" >
        <table className="table table-hover">
                <thead>
                    <tr className="table-info">
                    <th scope="col">Jugador</th>
                    <th scope="col">Juegos Ganados</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        datos.map( ( user ,i) => {
                            return (
                                <tr key={i}>
                                    <th scope="row">Jugador No. { user._id }</th>
                                    <th >{ user.count }</th>
                                </tr>
                            );
                        }) 
                    }
                    
                </tbody>
            </table>
    </div>
    )
}
