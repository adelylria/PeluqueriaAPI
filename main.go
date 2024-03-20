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
	log.Printf("Listening to port :8080")
	http.ListenAndServe(port, nil)

}
