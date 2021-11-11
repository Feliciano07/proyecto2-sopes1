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
	game := Juego{
		Request_number: 1,
		Game:           2,
		Gamename:       "Game1",
		Winner:         "001",
		Players:        20,
		Worker:         "Kafka",
	}
	salida := Parser(game)
	produce(salida)
}

func Parser(game Juego) string {
	b, err := json.Marshal(game)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(b))
	return string(b)
}

func produce(msg string) {

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:   []string{broker1Address, broker2Address, broker3Address},
		Topic:     topic,
		BatchSize: 1,
	})

	err := w.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte("Key-A"),
		Value: []byte(msg),
	})

	if err != nil {
		fmt.Println("Error message")
		panic("could not write message " + err.Error())
	}
	fmt.Println("writes message")
	time.Sleep(3 * time.Second)
}
