package apperrors

import "errors"

// ErrUserNotFound indica que el usuario no fue encontrado en la base de datos.
var ErrUserNotFound = errors.New("usuario no encontrado")

// ErrInvalidCredentials indica que las credenciales proporcionadas son inválidas.
var ErrInvalidCredentials = errors.New("credenciales inválidas")
