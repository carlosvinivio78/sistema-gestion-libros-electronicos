package main

import (
	"fmt"
	"sistema-gestion-libros-electronicos/models"
	"sistema-gestion-libros-electronicos/services"
)

func main() {
	libros := []models.Libro{}

	libro := models.Libro{
		ID:     1,
		Titulo: "Go Básico",
		Autor:  "Autor Ejemplo",
	}

	libros = services.AgregarLibro(libros, libro)

	fmt.Println("Libros registrados:")
	for _, l := range services.ListarLibros(libros) {
		fmt.Println(l.Titulo, "-", l.Autor)
	}
}
