package services

type Usuario struct {
	id     int
	nombre string
	email  string
}

func NuevoUsuario(id int, nombre, email string) *Usuario {
	return &Usuario{
		id:     id,
		nombre: nombre,
		email:  email,
	}
}

func (u *Usuario) GetID() int {
	return u.id
}

func (u *Usuario) GetNombre() string {
	return u.nombre
}

func (u *Usuario) GetEmail() string {
	return u.email
}
