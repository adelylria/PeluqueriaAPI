package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/login", loginHandler).Methods("POST")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

}
