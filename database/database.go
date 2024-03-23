package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() error {

	// Accede a las variables de entorno cargadas
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construye la cadena de conexión
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Println("Error de conexión a la base de datos:", err)
		return err
	}

	return nil
}

// GetDB devuelve la instancia de la base de datos
func GetDB() *sql.DB {
	return db
}
