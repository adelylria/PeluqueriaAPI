package routes

import (
	"net/http"

	"github.com/adelylria/PeluqueriaAPI/handlers"
	"github.com/gorilla/mux"
)

func RegisterAPIRoutes(r *mux.Router) {
	r.PathPrefix("/api/").Handler(http.StripPrefix("/api", http.FileServer(http.Dir("./web/api/"))))
	r.HandleFunc("/api/login", handlers.LoginHandler).Methods("POST")
	//r.HandleFunc("/api/admin", handlers).Methods("POST")
}

// RegisterWebRoutes registra las rutas para las páginas web estáticas
func RegisterWebRoutes(r *mux.Router) {
	// Rutas estáticas para las páginas de login y registro
	r.PathPrefix("/login/").Handler(http.StripPrefix("/login/", http.FileServer(http.Dir("./web/login/"))))
	r.PathPrefix("/register/").Handler(http.StripPrefix("/register/", http.FileServer(http.Dir("./web/register/"))))
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web"))))
}
