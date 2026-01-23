package services

import "sistema-gestion-libros-electronicos/models"

func AgregarLibro(libros []models.Libro, libro models.Libro) []models.Libro {
	return append(libros, libro)
}

func ListarLibros(libros []models.Libro) []models.Libro {
	return libros
}
