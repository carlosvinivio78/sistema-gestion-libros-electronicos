package services

import (
	"errors"
	"strings"

	"sistema-gestion-libros-electronicos/models"
)

// LibroServiceInterface define el contrato que debe cumplir el servicio.
// Permite aplicar el concepto de interfaces en Go.
type LibroServiceInterface interface {
	AgregarLibro(l *models.Libro) error
	ListarLibros() []*models.Libro
	BuscarLibroPorID(id int) (*models.Libro, error)
	BuscarPorTitulo(titulo string) []*models.Libro
	ActualizarLibro(id int, titulo, autor, categoria string, anio int) error
	EliminarLibro(id int) error
	ContarLibros() int
}

// LibroService contiene la lógica del sistema.
// Utiliza un map para almacenar libros dinámicamente.
type LibroService struct {
	libros map[int]*models.Libro
}

// NuevoLibroService inicializa el servicio con un map vacío.
func NuevoLibroService() *LibroService {
	return &LibroService{
		libros: make(map[int]*models.Libro),
	}
}

// AgregarLibro añade un libro al sistema.
// Verifica que no exista un libro con el mismo ID.
func (s *LibroService) AgregarLibro(l *models.Libro) error {

	if _, existe := s.libros[l.GetID()]; existe {
		return errors.New("el libro ya existe")
	}

	s.libros[l.GetID()] = l
	return nil
}

// ListarLibros devuelve todos los libros almacenados.
func (s *LibroService) ListarLibros() []*models.Libro {

	var lista []*models.Libro

	for _, libro := range s.libros {
		lista = append(lista, libro)
	}

	return lista
}

// BuscarLibroPorID busca un libro específico por su identificador.
func (s *LibroService) BuscarLibroPorID(id int) (*models.Libro, error) {

	libro, existe := s.libros[id]

	if !existe {
		return nil, errors.New("libro no encontrado")
	}

	return libro, nil
}

// BuscarPorTitulo permite buscar libros por coincidencia parcial de título.
// Se usa strings.ToLower para hacer búsqueda sin importar mayúsculas.
func (s *LibroService) BuscarPorTitulo(titulo string) []*models.Libro {

	var resultados []*models.Libro

	for _, libro := range s.libros {
		if strings.Contains(
			strings.ToLower(libro.GetTitulo()),
			strings.ToLower(titulo),
		) {
			resultados = append(resultados, libro)
		}
	}

	return resultados
}

// ActualizarLibro modifica un libro existente.
func (s *LibroService) ActualizarLibro(
	id int,
	titulo, autor, categoria string,
	anio int,
) error {

	libro, existe := s.libros[id]

	if !existe {
		return errors.New("libro no encontrado")
	}

	return libro.Actualizar(titulo, autor, categoria, anio)
}

// EliminarLibro elimina un libro del sistema.
func (s *LibroService) EliminarLibro(id int) error {

	if _, existe := s.libros[id]; !existe {
		return errors.New("no se puede eliminar, libro no existe")
	}

	delete(s.libros, id)
	return nil
}

// ContarLibros devuelve la cantidad total de libros registrados.
func (s *LibroService) ContarLibros() int {
	return len(s.libros)
}
