package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/adelylria/PeluqueriaAPI/apperrors"
	"github.com/adelylria/PeluqueriaAPI/database/repository"
	"github.com/adelylria/PeluqueriaAPI/models"
	"github.com/golang-jwt/jwt/v5"
)

// LoginHandler maneja la solicitud de inicio de sesión
// LoginHandler maneja la solicitud de inicio de sesión
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica si el método es POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Configura el encabezado Content-Type para aceptar JSON
	w.Header().Set("Content-Type", "application/json")

	// Decodifica el cuerpo JSON de la solicitud en una estructura de usuario
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Formato JSON inválido", http.StatusBadRequest)
		return
	}

	// Busca el usuario en la base de datos por nombre de usuario y contraseña
	userData, err := repository.GetUserByUsernameAndPassword(user.Username, user.Password)
	if err != nil {
		switch err {
		case apperrors.ErrUserNotFound:
			http.Error(w, "Credenciales incorrectas", http.StatusUnauthorized)
			log.Println("Usuario no encontrado en la base de datos")
		case apperrors.ErrInvalidCredentials:
			http.Error(w, "Credenciales incorrectas", http.StatusUnauthorized)
			log.Printf("Credenciales incorrectas para el usuario: %s", user.Username)
		default:
			http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
			log.Println("Error al consultar la base de datos:", err)
		}
		return
	}

	// Genera el token
	token, err := generateToken(userData.ID, user.Username)
	if err != nil {
		http.Error(w, "Error al generar el token", http.StatusInternalServerError)
		log.Println("Error al generar el token:", err)
		return
	}

	// Crea una estructura de respuesta
	loginResponse := models.LoginResponse{
		Message: "Inicio de sesión exitoso",
		Token:   token,
		User: models.LoginResponseUser{
			ID:    userData.ID,
			Email: userData.Email,
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

// generateToken genera un token de autenticación JWT para un usuario dado.
// Recibe el ID de usuario y el nombre de usuario como parámetros y devuelve el token generado y un error, si lo hay.
func generateToken(userID int, username string) (string, error) {
	// Obtenemos la llave secreta
	s_key := os.Getenv("SECRET_KEY")
	// Clave secreta para firmar el token
	secretKey := []byte(s_key)

	// Duración del token (En horas, se puede modificar)
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
