package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic          = "tema2"
	grupo          = "miGrupo"
	broker1Address = "34.125.238.189:19092"
	broker2Address = "34.125.238.189:29092"
	broker3Address = "34.125.238.189:39092"
)

type Juego struct {
	Request_number int    `json:"request_number"`
	Game           int    `json:"game"`
	Gamename       string `json:"gamename"`
	Winner         string `json:"winner"`
	Players        int    `json:"players"`
	Worker         string `json:"worker"`
}

func main() {
	consume()
}

func consume() {

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
		mensaje := ConvertJ(string(m.Value))
		fmt.Println(mensaje.Worker)
	}
}

func ConvertJ(msg string) Juego {

	var salida Juego
	err := json.Unmarshal([]byte(msg), &salida)
	if err != nil {
		fmt.Println("Error al convertir")
	}
	return salida
}
