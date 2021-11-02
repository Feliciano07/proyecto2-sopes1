package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

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

func ConvertJ(msg string) Juego {

	var salida Juego
	err := json.Unmarshal([]byte(msg), &salida)
	if err != nil {
		fmt.Println("Error al convertir")
	}
	return salida
}

func EnviarBase(msg string) {
	data := ConvertJ(msg)
	fmt.Println(data.Worker)
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
			EnviarBase(string(msg.Data))
			cancel()
		})
		if err != nil {
			fmt.Println("Error al recibir")
		}
	}
}

func main() {
	CargarClaves()
	Suscriber()
}
