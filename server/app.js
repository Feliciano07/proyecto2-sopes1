const express = require('express');
const cors = require('cors');
const {dbMongo} = require('./database/config');
const http = require('http');
const socketIo = require('socket.io');
const Juegos= require ('./models/juego');


const app = express();
app.set('port', process.env.PORT || 8000);
dbMongo();
app.use(cors());
app.use(express.json());

const server = http.createServer(app);

const io = socketIo(server, {
    cors:{
        origin:'*',
    },
});

let interval;
let interval2;
let interval3;
let interval4;

io.on('connection', socket=>{
    console.log("We have a new conecction!!");
    if (interval) {
        clearInterval(interval);
        clearInterval(interval2);
        clearInterval(interval3);
        clearInterval(interval4);
    }

    interval = setInterval(() => {
        Juegos.find().limit(50)
            .exec()
            .then( x =>  socket.emit("Juegos",x) );
    }, 8000);

    interval2 = setInterval(() => {
        Juegos.aggregate([
            {
                $sortByCount: '$gamename'
            }
        ])
        .limit(3)
        .exec()
        .then(x => socket.emit("top3", x));
    
      }, 5000);

    interval3 = setInterval(() => {
        Juegos.aggregate([
            {
                $sortByCount: '$worker'
            }
        ])
        .exec()
        .then(x => socket.emit('worker', x))
    
    }, 5000);

    
     interval4 = setInterval(()=>{
        Juegos.find().sort({_id:-1}).limit(10)
        .exec()
        .then(x=> socket.emit('lastGame', x));
     },5000)

    interval5 = setInterval(()=>{
        Juegos.aggregate([
            {
                $sortByCount: '$winner'
            }
        ])
        .limit(10)
        .exec()
        .then(x => socket.emit('mejores',x));
    },5000)

    socket.on("disconnect", () => {
        console.log("Client had left!!");
        clearInterval(interval);
        clearInterval(interval2);
        clearInterval(interval3);
        clearInterval(interval4);
    });
});




app.use("/", require('./routes/index.routes'));
server.listen( app.get('port') , () => console.log(`Listening on port ${ app.get('port') }`));