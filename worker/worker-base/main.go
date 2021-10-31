package main

import (
	"fmt"
	"net/http"
	"context"
	"log"
	"time"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func index(w http.ResponseWriter, r *http.Request) {
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
