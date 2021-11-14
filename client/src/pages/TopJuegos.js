import React,{useEffect,useState} from 'react'
import {CanvasJSChart} from 'canvasjs-react-charts'

export const TopJuegos = ({socket}) => {

    const [datos, setDatos] = useState([]);

    useEffect(() => {
        socket.on("top3", data =>  setDatos(data) );
      }, []);

    const grafica = datos.map( d => {
        return {label: d._id, y:d.count}
    })

    const optionsGame = {
        exportEnabled: true,
        animationEnabled: true,
        title: {
            text: "Top 3 Juegos"
        },
        data: [{
            type: "column",
            dataPoints: grafica
        }]
    }

    return (
        <div className="col-md-10 offset-1 mt-4 mb-4" >
            <CanvasJSChart options={ optionsGame } />
        </div>
    )
}
