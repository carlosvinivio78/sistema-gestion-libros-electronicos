/*
@Autor : Carlos Ron
@Fecha: 28/02/2026
@Descripcion : Desarollo de la aplicacion web "Sistema de gestion de libros electronicos"
*/
package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Variable global que almacenará la conexión a la base de datos
var DB *sql.DB

// Establece la conexión con la base de datos MySQL
func Connect() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error cargando .env: %v", err)
	}
	//Formato: usuario:contraseña@tcp(host:puerto)/nombre_base_datos
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	// Apertura de la conexión con MySQL
	var err2 error
	DB, err2 = sql.Open("mysql", dsn)
	if err2 != nil {
		return nil, err2
	}
	// Verificación de que la conexión sea válida
	if err2 = DB.Ping(); err2 != nil {
		return nil, err2
	}
	// Mensaje en consola indicando conexión exitosa
	log.Println(" Conexión exitosa con la BD")
	return DB, nil
}
