package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/adelylria/PeluqueriaAPI/database/repository"
	"github.com/golang-jwt/jwt/v5"
)

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica si el método es GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	// Valida el token de autorización
	token, err := validateToken(r)
	if err != nil {
		// Maneja el error de token inválido o faltante
		http.Error(w, "Token de autorización inválido", http.StatusUnauthorized)
		return
	}

	// Extrae el ID de usuario del token
	userID := int(token.Claims.(jwt.MapClaims)["user_id"].(float64))

	// Obtén la información del usuario por su ID
	user, err := repository.GetUserByID(userID)
	if err != nil {
		http.Error(w, "Error al obtener información del usuario", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	// Envía la información del usuario como respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// validateToken valida el token de autorización y devuelve los claims del token si es válido
func validateToken(r *http.Request) (*jwt.Token, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("token de autorización faltante")
	}

	// Verifica si el esquema del encabezado de autorización es "Bearer"
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		return nil, errors.New("token de autorización invalido")
	}

	// Extrae el token de autenticación
	tokenString := authParts[1]

	// Obtiene la clave secreta
	secretKey := []byte(os.Getenv("SECRET_KEY"))

	// Parsea y verifica el token con la clave secreta
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifica que el método de firma sea el esperado (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
		}
		return secretKey, nil

	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
