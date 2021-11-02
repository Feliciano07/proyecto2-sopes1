package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
)

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

func Parser(game Juego) string {
	b, err := json.Marshal(game)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(b))
	return string(b)
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

func publish(msg string) error {
	projectID := goDotEnvVariable("PROJECT_ID")
	topicID := goDotEnvVariable("TOPIC_ID")
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Println("error")
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{Data: []byte(msg)})

	id, err := result.Get(ctx)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Error: %v", err)
	}
	fmt.Println("Published a message; msg ID: %v\n", id)
	return nil
}

func main() {
	CargarClaves()
	game := Juego{
		Request_number: 1,
		Game:           2,
		Gamename:       "Game1",
		Winner:         "001",
		Players:        20,
		Worker:         "PubSub",
	}
	salida := Parser(game)
	publish(salida)
}
