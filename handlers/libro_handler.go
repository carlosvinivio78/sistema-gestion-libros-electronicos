/*
@Autor : Carlos Ron
@Fecha: 28/02/2026
@Descripcion : Desarollo de la aplicacion web "Sistema de gestion de libros electronicos"
*/
package handlers

import (
	"html/template"
	"net/http"
	"sistema-gestion-libros-electronicos/models"
	"strconv"
)

// LISTAR
func ListarLibros(w http.ResponseWriter, r *http.Request) {
	libros, _ := models.GetLibros()
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/lista.html"))
	tmpl.ExecuteTemplate(w, "base", libros)
}

// Agregar un nuevo libro.
func FormCrear(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/crear.html"))
	tmpl.ExecuteTemplate(w, "base", nil)
}

// GUARDAR
func GuardarLibro(w http.ResponseWriter, r *http.Request) {
	anio, _ := strconv.Atoi(r.FormValue("anio"))

	libro := models.Libro{
		Titulo:    r.FormValue("titulo"),
		Autor:     r.FormValue("autor"),
		Anio:      anio,
		Categoria: r.FormValue("categoria"),
	}

	models.CrearLibro(libro)
	http.Redirect(w, r, "/libros", http.StatusSeeOther)
}

// BUSCAR POR ID
func BuscarID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	libro, _ := models.GetLibroByID(id)

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/editar.html"))
	tmpl.ExecuteTemplate(w, "base", libro)
}

// BUSCAR POR TITULO
func BuscarTitulo(w http.ResponseWriter, r *http.Request) {
	titulo := r.FormValue("titulo")
	libros, _ := models.BuscarPorTitulo(titulo)

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/lista.html"))
	tmpl.ExecuteTemplate(w, "base", libros)
}

// ACTUALIZAR
func Actualizar(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	anio, _ := strconv.Atoi(r.FormValue("anio"))

	libro := models.Libro{
		ID:        id,
		Titulo:    r.FormValue("titulo"),
		Autor:     r.FormValue("autor"),
		Anio:      anio,
		Categoria: r.FormValue("categoria"),
	}

	models.ActualizarLibro(libro)
	http.Redirect(w, r, "/libros", http.StatusSeeOther)
}

// ELIMINAR
func Eliminar(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	models.EliminarLibro(id)
	http.Redirect(w, r, "/libros", http.StatusSeeOther)
}

// CONTAR
func Contar(w http.ResponseWriter, r *http.Request) {
	total, _ := models.ContarLibros()
	w.Write([]byte("Total de libros: " + strconv.Itoa(total)))
}

// SALIR
func Salir(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sistema cerrado correctamente"))
}
