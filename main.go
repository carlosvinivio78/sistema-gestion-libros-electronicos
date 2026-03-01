/*
@Autor : Carlos Ron
@Fecha: 28/02/2026
@Descripcion : Desarollo de la aplicacion web "Sistema de gestion de libros electronicos"
*/
package main

import (
	"log"
	"net/http"

	"sistema-gestion-libros-electronicos/db"
	"sistema-gestion-libros-electronicos/handlers"
)

func main() {

	//  Conectar a la base de datos y validar error
	_, err := db.Connect()
	if err != nil {
		log.Fatal("Error conectando a la BD:", err)

	}
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

	log.Println(" Servidor en http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
