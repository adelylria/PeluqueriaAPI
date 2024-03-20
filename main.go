package main

import (
	"log"
	"net/http"
	"os"

	"github.com/adelylria/PeluqueriaAPI/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	http.Handle("/", r)
	port := os.Getenv("PORT")
	done := make(chan bool)
	go http.ListenAndServe(port, nil)
	log.Printf("Listening to port %v", port)

	<-done
}
