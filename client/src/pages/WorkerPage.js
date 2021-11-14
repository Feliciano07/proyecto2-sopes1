import React,{useState, useEffect} from 'react'
import {CanvasJSChart} from 'canvasjs-react-charts'

export const WorkerPage = ({socket}) => {

    const [datos, setDatos] = useState([]);

    useEffect(() => {
        socket.on("worker", data =>  setDatos(data) );
    }, [datos]);

    const grafica = datos.map( d => {
        return {label: d._id, y:d.count}
    })

    const optionsGame = {
        exportEnabled: true,
        animationEnabled: true,
        title: {
            text: "Comparacion de Worker"
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
