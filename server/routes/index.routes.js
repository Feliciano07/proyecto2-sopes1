const express = require('express');
const {response,request} = require('express');
const router = express.Router();
const Juegos= require ('../models/juego');
// const key = "lista"

// const redis = require('redis');
// const client = redis.createClient({
//     port:6379, 
//     host:"34.125.72.176",
//     db:0
// })


router.get('/', (req = request, res = response)=>{
    res.send("Server corriendo")
});

// REPORTES DE MONGODB

/*
 - Reporte Datos almacenados
 - Tabla con los logs almacenados
*/
router.get('/juegos', (req = request, res = response)=>{
    Juegos.find()
        .exec()
        .then(x => res.status(200).send(x));
});

/**
 * Grafica del top 3 juegos
 */
router.get('/top-juegos', (req= request, res = response)=>{
    Juegos.aggregate([
        {
            $sortByCount: '$gamename'
        }
    ])
    .limit(3)
    .exec()
    .then(x => res.status(200).send(x));
});

/**
 * Grafica que compara a los 3 workes
 */
router.get('/worker', (req = request, res = response)=>{
    Juegos.aggregate([
        {
            $sortByCount: '$worker'
        }
    ])
    .exec()
    .then(x => res.status(200).send(x));
})

/**
 * Ultimos 10 juegos
 */
router.get('/lastGame', async (req=request, res= response) =>{
    Juegos.find().sort({_id:-1}).limit(10)
    .exec()
    .then(x=> res.status(200).send(x));
})

/**
 * Los 10 mejores jugadores
 */
router.get('/top-jugadores', (req = request, res=response)=>{
    Juegos.aggregate([
        {
            $sortByCount: '$winner'
        }
    ])
    .limit(10)
    .exec()
    .then(x => res.status(200).send(x));
})


// const SetearDatos = async() =>{
//     const data = await Juegos.find();
//     const json = JSON.stringify(data);
//     client.set(key,json, (err, result) =>{
//         if(err){
//             console.log(err);
//         }
//         console.log(result);
//         return true
//     })
// }


module.exports = router;