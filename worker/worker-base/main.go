package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type Juego struct {
	ID      primitive.ObjectID `bson:"_id"`
	Mensaje string             `bson:"text"`
}

type JuegoRedis struct {
	Mensaje string
}

const MongoDb = "mongodb://grupo33:pass%2B1234@34.125.72.176:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"

func GuardarMongo(mensaje string) {
	//Conexion a mongo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoclient, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}
	Database := mongoclient.Database("db_sopes")
	Collection := Database.Collection("data")

	juego := &Juego{
		ID:      primitive.NewObjectID(),
		Mensaje: "hola",
	}

	infectadosResult, err := Collection.InsertOne(ctx, juego)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(infectadosResult)
}

func GuardarRedis(mensaje string) {

	client := redis.NewClient(&redis.Options{
		Addr:     "34.125.72.176:6379",
		DB:       1,
		Password: "",
	})
	defer client.Close()

	game := JuegoRedis{
		Mensaje: "hola",
	}

	//Guardar en redis
	val, err := client.Do(ctx, "RPUSH", "lista", game).Result()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(val)
}

func (ac JuegoRedis) MarshalBinary() ([]byte, error) {
	return json.Marshal(ac)
}

func index(w http.ResponseWriter, r *http.Request) {
	GuardarRedis(`{Mensaje: "hola"}`)
	fmt.Fprintf(w, "API esta funcionando")
}

func main() {

	//Rutas del servidor
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index).Methods("GET")

	//Cors
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	http.ListenAndServe(":3002", handler)
}
