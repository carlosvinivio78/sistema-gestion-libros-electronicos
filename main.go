package main

import (
	"log"
	"net/http"
	"sistema-gestion-libros-electronicos/db"
	"sistema-gestion-libros-electronicos/handlers"
	"sistema-gestion-libros-electronicos/routes"

	_ "github.com/gin-gonic/gin"
)

func main() {
	// Conectar a la BD
	_, err := db.Connect()
	if err != nil {
		log.Fatal("Error conectando a la BD:", err)
	}

	// --- Servidor HTML existente ---
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/libros", handlers.ListarLibros)
	http.HandleFunc("/libros/nuevo", handlers.FormCrear)
	http.HandleFunc("/libros/guardar", handlers.GuardarLibro)
	http.HandleFunc("/libros/buscarID", handlers.BuscarID)
	http.HandleFunc("/libros/buscarTitulo", handlers.BuscarTitulo)
	http.HandleFunc("/libros/actualizar", handlers.Actualizar)
	http.HandleFunc("/libros/eliminar", handlers.Eliminar)
	http.HandleFunc("/libros/contar", handlers.Contar)
	http.HandleFunc("/salir", handlers.Salir)

	go func() {
		log.Println("Servidor HTML en http://localhost:8000")
		log.Fatal(http.ListenAndServe(":8000", nil))
	}()

	// --- Servidor API JSON con Gin ---
	r := routes.SetupRouter()
	log.Println("API JSON en http://localhost:8080")
	r.Run(":8080")
}
