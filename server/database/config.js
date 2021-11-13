
const mongoose = require('mongoose');

const dbMongo = async () =>{
    try {
        await mongoose.connect(
            "mongodb://grupo33:pass%2B1234@34.125.72.176:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false",
            {
            useNewUrlParser: true,
            useUnifiedTopology: true,
            dbName:'db_sopes'
            });
        console.log('Base de datos conectada');
    } catch (error) {
        console.log(error);
        throw new Error('Error en la conexion con la base de datos');
    }
}

module.exports = {
    dbMongo
}