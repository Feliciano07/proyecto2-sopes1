package main

import (
	//Propias
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	//Database
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//PubSub
	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
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

func CargarClaves() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./clave.json")
}

func goDotEnvVariable(key string) string {

	// Leer el archivo .env ubicado en la carpeta actual
	err := godotenv.Load(".env")

	// Si existio error leyendo el archivo
	if err != nil {
		log.Fatalf("Error cargando las variables de entorno")
	}

	// Enviar la variable de entorno que se necesita leer
	return os.Getenv(key)
}

func Convert_string_to_Juego(msg string) Juego {

	var salida Juego
	err := json.Unmarshal([]byte(msg), &salida)
	if err != nil {
		fmt.Println("Error al convertir")
	}
	return salida
}

func Suscriber() error {
	projectID := goDotEnvVariable("PROJECT_ID")
	SubID := goDotEnvVariable("TOPIC_ID")
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Println("error")
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	var mu sync.Mutex
	for {
		sub := client.Subscription(SubID)
		cctx, cancel := context.WithCancel(ctx)
		err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
			mu.Lock()
			defer mu.Unlock()
			msg.Ack()
			Save_database(string(msg.Data))
			cancel()
		})
		if err != nil {
			fmt.Println("Error al recibir")
		}
	}
}

//Esta es la funcion que se encarga de obtener el mensaje y poder guardar en las bases
func Save_database(mensaje string) {
	data := Convert_string_to_Juego(mensaje)
	GuardarMongo(data)
	GuardarRedis(data)
}

func main() {
	CargarClaves()
	Suscriber()
}
