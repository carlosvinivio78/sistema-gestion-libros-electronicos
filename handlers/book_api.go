package handlers

import (
	"net/http"
	"strconv"

	"sistema-gestion-libros-electronicos/models"

	"github.com/gin-gonic/gin"
)

// Listar libros en JSON
func GetBooksJSON(c *gin.Context) {
	libros, _ := models.GetLibros()
	c.JSON(http.StatusOK, libros)
}

// Obtener libro por ID en JSON
func GetBookByIDJSON(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	libro, err := models.GetLibroByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Libro no encontrado"})
		return
	}
	c.JSON(http.StatusOK, libro)
}

// Crear libro desde JSON
func CreateBookJSON(c *gin.Context) {
	var libro models.Libro
	if err := c.ShouldBindJSON(&libro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.CrearLibro(libro)
	c.JSON(http.StatusCreated, libro)
}

// Actualizar libro desde JSON
func UpdateBookJSON(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var libro models.Libro
	if err := c.ShouldBindJSON(&libro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	libro.ID = id
	models.ActualizarLibro(libro)
	c.JSON(http.StatusOK, libro)
}

// Eliminar libro desde JSON
func DeleteBookJSON(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	models.EliminarLibro(id)
	c.JSON(http.StatusOK, gin.H{"message": "Libro eliminado"})
}
