const express = require('express');
const {response,request} = require('express');
const redis = require('redis');


const Juegos= require ('../models/juego');

const router = express.Router();

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

// REPORTES DE REDI


module.exports = router;