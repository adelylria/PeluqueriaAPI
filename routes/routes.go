package routes

import (
	"github.com/adelylria/PeluqueriaAPI/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/api/admin", handlers.HairdresserLogin).Methods("POST")
}
