package models

import "errors"

// Libro representa la estructura principal del sistema.
// Sus atributos son privados (minúscula) para aplicar encapsulación.
type Libro struct {
	id        int
	titulo    string
	autor     string
	categoria string
	anio      int
}

// NuevoLibro funciona como constructor.
// Permite crear un nuevo libro con sus datos iniciales.
func NuevoLibro(id int, titulo, autor, categoria string, anio int) *Libro {
	return &Libro{
		id:        id,
		titulo:    titulo,
		autor:     autor,
		categoria: categoria,
		anio:      anio,
	}
}

// Métodos GET para acceder a los datos (encapsulación).
func (l *Libro) GetID() int {
	return l.id
}

func (l *Libro) GetTitulo() string {
	return l.titulo
}

func (l *Libro) GetAutor() string {
	return l.autor
}

func (l *Libro) GetCategoria() string {
	return l.categoria
}

func (l *Libro) GetAnio() int {
	return l.anio
}

// Actualizar modifica los datos del libro.
// Valida que el título y autor no estén vacíos.
func (l *Libro) Actualizar(titulo, autor, categoria string, anio int) error {

	if titulo == "" || autor == "" {
		return errors.New("titulo y autor no pueden estar vacíos")
	}

	l.titulo = titulo
	l.autor = autor
	l.categoria = categoria
	l.anio = anio

	return nil
}
