package main

import (
	"fmt"
	"sistema-gestion-libros-electronicos/models"
	"sistema-gestion-libros-electronicos/services"
)

func main() {
	libros := []models.Libro{}

	libro1 := models.Libro{
		ID:     1,
		Titulo: "A Sangre Fria ",
		Autor:  "Truman Capote",
	}
	libro2 := models.Libro{
		ID:     2,
		Titulo: "Cien Años de Soledad ",
		Autor:  "Gabriel García Márquez",
	}
	libro3 := models.Libro{
		ID:     3,
		Titulo: "El Principito ",
		Autor:  "Antoine de Saint-Exupéry",
	}

	libros = services.AgregarLibro(libros, libro1)
	libros = services.AgregarLibro(libros, libro2)
	libros = services.AgregarLibro(libros, libro3)

	fmt.Println("Libros registrados:")
	for _, l := range services.ListarLibros(libros) {
		fmt.Println(l.Titulo, "-", l.Autor)
	}
}
