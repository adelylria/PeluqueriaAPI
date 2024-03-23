package main

import (
	"log"
	"net/http"
	"os"

	"github.com/adelylria/PeluqueriaAPI/database"
	"github.com/adelylria/PeluqueriaAPI/routes"
	"github.com/gorilla/mux"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatal("Error al inicializar la base de datos:", err)
	}

	r := mux.NewRouter()

	routes.RegisterAPIRoutes(r)
	routes.RegisterWebRoutes(r)

	http.Handle("/", r)

	port := os.Getenv("PORT")
	log.Print("Listening to port :8080")
	http.ListenAndServe(port, nil)

}
