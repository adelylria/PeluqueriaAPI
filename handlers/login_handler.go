package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/adelylria/PeluqueriaAPI/database"
	"github.com/adelylria/PeluqueriaAPI/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Message string                   `json:"message"`
	Token   string                   `json:"token"`
	User    models.LoginResponseUser `json:"user"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica si el método es POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Configura el encabezado Content-Type para aceptar JSON
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Formato JSON inválido", http.StatusBadRequest)
		return
	}

	// Obtiene la instancia de la base de datos
	db := database.GetDB()

	// Realiza una consulta para obtener el hash de contraseña y el ID del usuario
	var hashedPassword, email string
	var userID int
	err := db.QueryRow("SELECT idusuario, email, password FROM usuario WHERE username = ?", user.Username).Scan(&userID, &email, &hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Credenciales incorrectas", http.StatusUnauthorized)
			log.Println("Usuario no encontrado en la base de datos")
			return
		}
		http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
		log.Println("Error al consultar la base de datos:", err)
		return
	}

	// Compara la contraseña ingresada con el hash almacenado en la base de datos
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)); err != nil {
		http.Error(w, "Credenciales incorrectas", http.StatusUnauthorized)
		log.Printf("Cliente: %s | Método: %s", r.RemoteAddr, r.Method)
		log.Printf("Credenciales incorrectas para el usuario: %s", user.Username)
		return
	}

	// Genera el token
	token, err := generateToken(userID, user.Username)
	if err != nil {
		http.Error(w, "Error al generar el token", http.StatusInternalServerError)
		log.Println("Error al generar el token:", err)
		return
	}

	// Crea una estructura de respuesta
	loginResponse := LoginResponse{
		Message: "Inicio de sesión exitoso",
		Token:   token,
		User: models.LoginResponseUser{
			ID:    userID,
			Email: email,
		},
	}

	// Codifica la respuesta como JSON
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(loginResponse); err != nil {
		http.Error(w, "Error al escribir la respuesta JSON", http.StatusInternalServerError)
		log.Println("Error al escribir la respuesta JSON:", err)
	}

	// Registra información de inicio de sesión
	log.Printf("Cliente: %s | Método: %s", r.RemoteAddr, r.Method)
	log.Printf("Inicio de sesión exitoso. Usuario: %s", user.Username)
}

func generateToken(userID int, username string) (string, error) {
	// obtenemos la llave secreta
	s_key := os.Getenv("SECRET_KEY")
	// Clave secreta para firmar el token
	secretKey := []byte(s_key)

	// Duración del token (en este caso, 6 horas)
	expirationTime := time.Now().Add(6 * time.Hour)
	// Crea las afirmaciones (claims) del token
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      expirationTime.Unix(), // Tiempo de expiración del token
	}

	// Crea el token con las afirmaciones y firma
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
