const express = require('express');
const cors = require('cors');
const {dbMongo} = require('../database/config');

class Server{


    constructor(){
        this.app = express();
        this.port = process.env.PORT || 8000;

        this.database();
        this.middleware();
        this.routes();
    }

    routes(){
        this.app.use("/", require('../routes/index.routes'));
    }

    middleware(){
        this.app.use(cors());
        this.app.use(express.json());
    }

    listen(){
        this.app.listen(this.port, () =>{
            console.log('Server start int port', this.port)
        })
    }

    async database(){
        await dbMongo();
    }

}

module.exports = Server;