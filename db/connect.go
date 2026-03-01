package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error cargando .env: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err2 error
	DB, err2 = sql.Open("mysql", dsn)
	if err2 != nil {
		return nil, err2
	}

	if err2 = DB.Ping(); err2 != nil {
		return nil, err2
	}

	log.Println("✅ Conexión exitosa con la BD")
	return DB, nil
}
