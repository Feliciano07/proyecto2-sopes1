const {Schema, model} = require('mongoose');

const JuegosSchema = Schema ({
    request_number: Number,
    game: Number,
    gamename: String,
    winner: String,
    worker: String
})

module.exports = model("juegos",JuegosSchema)