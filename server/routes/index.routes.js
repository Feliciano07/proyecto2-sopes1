const express = require('express');
const {response,request} = require('express');
const router = express.Router();
const Juegos= require ('../models/juego');

const redis = require('redis');
const client = redis.createClient({
    port:6379, host:"34.125.245.149", db:1
})


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

// REPORTE DE REDIS

router.get('/prueba', (req = request, res = response) =>{
    client.set('lista', "valores", redis.print);
    res.status(200).send("Agregado")
})

module.exports = router;