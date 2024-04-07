package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adelylria/PeluqueriaAPI/database"
	"github.com/adelylria/PeluqueriaAPI/models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Message: "Formato JSON inválido"})
		return
	}

	// Genera un hash de la contraseña del usuario
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Message: "Error al encriptar la contraseña"})
		log.Println("Error al encriptar la contraseña:", err)
		return
	}

	// Obtén la instancia de la base de datos
	db := database.GetDB()

	isAdmin := 0
	if user.Admin {
		isAdmin = 1
	} else {
		isAdmin = 0
	}

	log.Printf("usernme: %s, password: %s, email: %s, admin: %v \n", user.Username, hashedPassword, user.Email, isAdmin)
	// Inserta el nuevo usuario en la base de datos con la contraseña encriptada
	_, err = db.Exec("INSERT INTO usuario (username, password, email, admin) VALUES (?, ?, ?, ?)", user.Username, hashedPassword, user.Email, isAdmin)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Message: "Error al registrar el usuario"})
		log.Println("Error al registrar el usuario en la base de datos:", err)
		return
	}

	// Responde con un mensaje de éxito
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario registrado con éxito"})
	log.Printf("Usuario registrado: %s", user.Username)
}
