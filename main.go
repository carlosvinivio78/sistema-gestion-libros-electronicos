// Autor: Carlos Ron
// Fecha: 14/02/2026
// Proyecto: Sistema de Gestión de Libros

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"sistema-gestion-libros-electronicos/models"
	"sistema-gestion-libros-electronicos/services"
)

func main() {

	// Se crea el servicio principal del sistema
	libroService := services.NuevoLibroService()

	// Reader permite leer datos desde teclado
	reader := bufio.NewReader(os.Stdin)

	for {

		// Menú principal del sistema
		fmt.Println("\n===== SISTEMA DE GESTIÓN DE LIBROS =====")
		fmt.Println("1. Agregar libro")
		fmt.Println("2. Listar libros")
		fmt.Println("3. Buscar libro por ID")
		fmt.Println("4. Buscar libro por título")
		fmt.Println("5. Actualizar libro")
		fmt.Println("6. Eliminar libro")
		fmt.Println("7. Contar libros")
		fmt.Println("8. Salir")
		fmt.Print("Seleccione una opción: ")

		// Se lee la opción ingresada
		opcionStr, _ := reader.ReadString('\n')
		opcionStr = strings.TrimSpace(opcionStr)
		opcion, _ := strconv.Atoi(opcionStr)

		switch opcion {

		case 1:
			// Registro de nuevo libro
			fmt.Print("ID: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))

			fmt.Print("Título: ")
			titulo, _ := reader.ReadString('\n')

			fmt.Print("Autor: ")
			autor, _ := reader.ReadString('\n')

			fmt.Print("Categoría: ")
			categoria, _ := reader.ReadString('\n')

			fmt.Print("Año: ")
			anioStr, _ := reader.ReadString('\n')
			anio, _ := strconv.Atoi(strings.TrimSpace(anioStr))

			libro := models.NuevoLibro(
				id,
				strings.TrimSpace(titulo),
				strings.TrimSpace(autor),
				strings.TrimSpace(categoria),
				anio,
			)

			err := libroService.AgregarLibro(libro)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Libro agregado correctamente.")
			}

		case 2:
			// Muestra todos los libros
			libros := libroService.ListarLibros()

			for _, libro := range libros {
				fmt.Println(
					libro.GetID(),
					"-",
					libro.GetTitulo(),
					"-",
					libro.GetAutor(),
				)
			}

		case 3:
			// Buscar libro por ID
			fmt.Print("Ingrese ID: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))

			libro, err := libroService.BuscarLibroPorID(id)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Encontrado:", libro.GetTitulo())
			}

		case 4:
			// Buscar libro por coincidencia de título
			fmt.Print("Ingrese título a buscar: ")
			titulo, _ := reader.ReadString('\n')

			resultados := libroService.BuscarPorTitulo(strings.TrimSpace(titulo))

			for _, libro := range resultados {
				fmt.Println(libro.GetID(), "-", libro.GetTitulo())
			}

		case 5:
			// Actualizar libro existente
			fmt.Print("ID del libro a actualizar: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))

			fmt.Print("Nuevo título: ")
			titulo, _ := reader.ReadString('\n')

			fmt.Print("Nuevo autor: ")
			autor, _ := reader.ReadString('\n')

			fmt.Print("Nueva categoría: ")
			categoria, _ := reader.ReadString('\n')

			fmt.Print("Nuevo año: ")
			anioStr, _ := reader.ReadString('\n')
			anio, _ := strconv.Atoi(strings.TrimSpace(anioStr))

			err := libroService.ActualizarLibro(
				id,
				strings.TrimSpace(titulo),
				strings.TrimSpace(autor),
				strings.TrimSpace(categoria),
				anio,
			)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Libro actualizado correctamente.")
			}

		case 6:
			// Eliminar libro
			fmt.Print("ID del libro a eliminar: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))

			err := libroService.EliminarLibro(id)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Libro eliminado.")
			}

		case 7:
			// Mostrar total de libros registrados
			fmt.Println("Total de libros:", libroService.ContarLibros())

		case 8:
			fmt.Println("Saliendo del sistema...")
			return

		default:
			fmt.Println("Opción inválida.")
		}
	}
}
