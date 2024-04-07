package main

import (
	"log"
	"net/http"
	"os"
	"time"

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
	port := os.Getenv("PORT")

	// Server configuration
	srv := &http.Server{
		Handler:      r,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	//http.Handle("/", r)

	log.Print("Listening to port :8080")
	srv.ListenAndServe()
}
