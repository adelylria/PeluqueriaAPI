package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Admin    bool   `json:"admin"`
}

type LoginResponseUser struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}
