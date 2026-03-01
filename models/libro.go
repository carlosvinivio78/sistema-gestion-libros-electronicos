/*
@Autor : Carlos Ron
@Fecha: 28/02/2026
@Descripcion : Desarollo de la aplicacion web "Sistema de gestion de libros electronicos"
*/
package models

import (
	"sistema-gestion-libros-electronicos/db"
)

// Atributos
type Libro struct {
	ID        int
	Titulo    string
	Autor     string
	Anio      int
	Categoria string
}

// AGREGAR
func CrearLibro(libro Libro) error {
	_, err := db.DB.Exec(
		"INSERT INTO libros(titulo, autor, anio, categoria) VALUES(?,?,?,?)",
		libro.Titulo, libro.Autor, libro.Anio, libro.Categoria)
	return err
}

// LISTAR TODOS
func GetLibros() ([]Libro, error) {
	rows, err := db.DB.Query("SELECT * FROM libros")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var libros []Libro

	for rows.Next() {
		var l Libro
		rows.Scan(&l.ID, &l.Titulo, &l.Autor, &l.Anio, &l.Categoria)
		libros = append(libros, l)
	}
	return libros, nil
}

// BUSCAR POR ID
func GetLibroByID(id int) (Libro, error) {
	var l Libro
	err := db.DB.QueryRow("SELECT * FROM libros WHERE id = ?", id).
		Scan(&l.ID, &l.Titulo, &l.Autor, &l.Anio, &l.Categoria)
	return l, err
}

// BUSCAR POR TITULO
func BuscarPorTitulo(titulo string) ([]Libro, error) {
	rows, err := db.DB.Query("SELECT * FROM libros WHERE titulo LIKE ?", "%"+titulo+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var libros []Libro
	for rows.Next() {
		var l Libro
		rows.Scan(&l.ID, &l.Titulo, &l.Autor, &l.Anio, &l.Categoria)
		libros = append(libros, l)
	}
	return libros, nil
}

// ACTUALIZAR
func ActualizarLibro(libro Libro) error {
	_, err := db.DB.Exec(
		"UPDATE libros SET titulo=?, autor=?, anio=?, categoria=? WHERE id=?",
		libro.Titulo, libro.Autor, libro.Anio, libro.Categoria, libro.ID)
	return err
}

// ELIMINAR
func EliminarLibro(id int) error {
	_, err := db.DB.Exec("DELETE FROM libros WHERE id=?", id)
	return err
}

// CONTAR
func ContarLibros() (int, error) {
	var total int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM libros").Scan(&total)
	return total, err
}
