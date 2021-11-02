package main

import (
	//Propias
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	//Database
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/segmentio/kafka-go"
)

const MongoDb = "mongodb://grupo33:pass%2B1234@34.125.72.176:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"

var ctx = context.Background()

type DataMongo struct {
	ID             primitive.ObjectID `bson:"_id"`
	Request_Number int                `bson:"request_number"`
	Game           int                `bson:"game"`
	GameName       string             `bson:"gamename"`
	Winner         string             `bson:"winner"`
	Players        int                `bson:"players"`
	Worker         string             `bson:"worker"`
}

func GuardarMongo(game Juego) {
	//Conexion a mongo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoclient, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}
	Database := mongoclient.Database("db_sopes")
	Collection := Database.Collection("juegos")

	juego := &DataMongo{
		ID:             primitive.NewObjectID(),
		Request_Number: game.Request_number,
		Game:           game.Game,
		GameName:       game.Gamename,
		Winner:         game.Winner,
		Players:        game.Players,
		Worker:         game.Worker,
	}

	infectadosResult, err := Collection.InsertOne(ctx, juego)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(infectadosResult)
}

func GuardarRedis(game Juego) {

	client := redis.NewClient(&redis.Options{
		Addr:     "34.125.72.176:6379",
		DB:       1,
		Password: "",
	})
	defer client.Close()

	//Guardar en redis
	val, err := client.Do(ctx, "RPUSH", "lista", game).Result()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(val)
}

func (ac Juego) MarshalBinary() ([]byte, error) {
	return json.Marshal(ac)
}

type Juego struct {
	Request_number int    `json:"request_number"`
	Game           int    `json:"game"`
	Gamename       string `json:"gamename"`
	Winner         string `json:"winner"`
	Players        int    `json:"players"`
	Worker         string `json:"worker"`
}

func Convert_string_to_Juego(msg string) Juego {

	var salida Juego
	err := json.Unmarshal([]byte(msg), &salida)
	if err != nil {
		fmt.Println("Error al convertir")
	}
	return salida
}

//Esta es la funcion que se encarga de obtener el mensaje y poder guardar en las bases
func Save_database(mensaje string) {
	data := Convert_string_to_Juego(mensaje)
	GuardarMongo(data)
	GuardarRedis(data)
}

//Kafka
const (
	topic          = "tema2"
	grupo          = "miGrupo"
	broker1Address = "34.125.238.189:19092"
	broker2Address = "34.125.238.189:29092"
	broker3Address = "34.125.238.189:39092"
)

func ConsumerKafka() {

	//l := log.New(os.Stdout, "kafka reader: ", 0)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address, broker2Address, broker3Address},
		Topic:   topic,
		MaxWait: 3 * time.Second,
		GroupID: grupo,
	})
	for {
		m, err := r.FetchMessage(context.Background())
		if err != nil {
			break
		}
		if err := r.CommitMessages(context.Background(), m); err != nil {
			fmt.Println("Error al confirmar")
		}
		Save_database(string(m.Value))
	}
}

func main() {
	ConsumerKafka()
}
