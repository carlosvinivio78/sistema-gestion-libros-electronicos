Sistema de Gestión de Libros Electrónicos

Datos del Grupo
Integrante: Carlos Ron  
Módulo: Desarrollo de Aplicaciones Web  

Objetivo del Programa
Desarrollar una aplicación web completa en lenguaje Go que implemente operaciones CRUD para la gestión de libros electrónicos, utilizando arquitectura por capas y base de datos MySQL. El sistema permite registrar, consultar, modificar y eliminar libros digitales de manera eficiente.

Fecha de entrega: 29 de Febrero de 2026

Funcionalidades:

El Sistema de Gestión de Libros Electrónicos es una aplicación web desarrollada en Go que permite administrar libros digitales mediante operaciones CRUD (Crear, Leer, Actualizar y Eliminar) conectadas a una base de datos MySQL.

El sistema fue desarrollado aplicando arquitectura por capas para mantener una estructura organizada y escalable.



- Agregar libro

- Buscar libro por ID

- Buscar libro por título

- Actualizar libro

- Eliminar libro

- Contar total de libros registrados

- Listar todos los libros


Tecnologías Utilizadas:

- Go (Backend)

- MySQL (Base de Datos)

- Laragon (Servidor Local)

- HTML (Interfaz)

- CSS (Diseño)

Arquitectura por Capas

Estructura del Proyecto:
sistema-libros-electronico/
├── cmd/
│ └── main.go
├── internal/
│ ├── handlers/
│ ├── models/
│ ├── repository/
│ └── services/
├── static/
│ ├── css/
│ └── js/
├── templates/
│ ├── index.html
│ ├── add.html
│ ├── edit.html
│ └── book.html
├── database/
│ └── schema.sql
├── go.mod
├── go.sum
└── README.md