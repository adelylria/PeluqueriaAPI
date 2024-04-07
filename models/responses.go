package models

type ErrorResponse struct {
	Message string `json:"message"`
}

type LoginResponseUser struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type LoginResponse struct {
	Message string            `json:"message"`
	Token   string            `json:"token"`
	User    LoginResponseUser `json:"user"`
}
