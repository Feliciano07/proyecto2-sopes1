const express = require('express');
const {response,request} = require('express');
const router = express.Router();
const Juegos= require ('../models/juego');

router.get('/', (req = request, res = response)=>{
    res.send("Server corriendo")
});


router.get('/listaJuegos', (req = request, res = response)=>{

    Juegos.find()
        .exec()
        .then(x => res.status(200).send(x));

})

module.exports = router;