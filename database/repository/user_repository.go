package repository

import (
	"database/sql"

	"github.com/adelylria/PeluqueriaAPI/apperrors"
	"github.com/adelylria/PeluqueriaAPI/database"
	"github.com/adelylria/PeluqueriaAPI/models"
	"golang.org/x/crypto/bcrypt"
)

// GetUserByID recupera un usuario de la base de datos por su ID.
// Recibe como parámetro el ID del usuario a recuperar.
// Devuelve un puntero al usuario y un posible error.
func GetUserByID(userID int) (*models.User, error) {
	var user models.User
	db := database.GetDB()

	err := db.QueryRow("SELECT idusuario, username, email FROM usuario WHERE idusuario = ?", userID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err != sql.ErrNoRows {
			// Otro error que no sea sql.ErrNoRows
			return nil, err
		}
		// Usuario no encontrado
		return nil, apperrors.ErrUserNotFound
	}

	return &user, nil
}

// GetUserByUsernameAndPassword busca un usuario en la base de datos por su nombre de usuario y contraseña.
// Devuelve los datos del usuario si las credenciales son válidas, de lo contrario, devuelve un error.
func GetUserByUsernameAndPassword(username, password string) (*models.User, error) {
	var user models.User

	// Obtiene la instancia de la base de datos
	db := database.GetDB()

	// Realiza una consulta para obtener los datos del usuario por su nombre de usuario
	err := db.QueryRow("SELECT idusuario, email, password FROM usuario WHERE username = ?", username).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrUserNotFound // Utiliza el error personalizado ErrUserNotFound
		}
		return nil, err // Otro error
	}

	// Compara la contraseña ingresada con la contraseña hash almacenada en la base de datos
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, apperrors.ErrInvalidCredentials // Utiliza el error personalizado ErrInvalidCredentials
	}

	return &user, nil // Credenciales válidas, devuelve los datos del usuario
}
