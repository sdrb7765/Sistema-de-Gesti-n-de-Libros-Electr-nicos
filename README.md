# Sistema-de-Gestion-de-Libros-Electronicos
Sistema de Gestión de Libros Electrónicos
"Estructura inicial del SGLE"
<img width="975" height="551" alt="image" src="https://github.com/user-attachments/assets/3a486007-3479-4142-8d0e-40b7305d65c5" />

auth
libros
catalogo
biblioteca
lector
README.md
main.go


# 📚 Librería Virtual en Go

Este proyecto es un avance del sistema **Librería Virtual**, desarrollado en el lenguaje Go como parte del curso de Programación.  
Incluye los contenidos esenciales de las **Unidades 1, 2 y 3**, aplicando funciones, estructuras de datos, orientación a objetos, encapsulación e interfaces.

---

## 🚀 Objetivo del Proyecto
El objetivo es simular un sistema básico de gestión de librería que permita:

- Registrar usuarios  
- Agregar y listar libros  
- Validar disponibilidad  
- Registrar préstamos  
- Actualizar el estado de los libros  

Este es el primer avance funcional del software.

---

##  Contenidos Aplicados (Requisitos del curso)

###  Unidad 1 – Fundamentos de Go
- Uso de funciones y parámetros  
- Uso de condicionales `if`  
- Ciclos `for range`  
- Paquetes e importaciones  
- Estructura básica `main.go`  

###  Unidad 2 – Estructuras de Datos
- Uso de **maps** para almacenar libros y usuarios  
- Uso de **slices** para almacenar préstamos  
- Uso de **structs** para representar entidades  
- Métodos y constructores personalizados  

### Unidad 3 – Programación Orientada a Objetos en Go
- Encapsulación mediante servicios  
- Interfaces (interface `Accion`)  
- Manejo de errores con `errors.New()`  
- Organización modular y profesional del proyecto  

---

## Estructura del Proyecto
# go.mod

module libreria-virtual

go 1.22

# main.go
package main

import (
	"fmt"
	"libreria-virtual/internal/models"
	"libreria-virtual/internal/services"
)

func main() {

	libroService := services.NuevoLibroService()
	usuarioService := services.NuevoUsuarioService()
	prestamoService := services.NuevoPrestamoService(libroService, usuarioService)

	// Crear usuarios
	usuario := models.NuevoUsuario(1, "Samuel Riera")
	usuarioService.AgregarUsuario(usuario)

	// Crear libros
	libro := models.NuevoLibro(101, "El señor de los anillos", "Tolkien")
	libroService.AgregarLibro(libro)

	// Préstamo
	err := prestamoService.CrearPrestamo(1, 101)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Préstamo realizado con éxito")
	}

	fmt.Println("Libros:", libroService.ListarLibros())
}

1. MODELOS (Structs + Constructores + Métodos)
# internal/models/libro.go

internal/models/libro.go
package models

// Libro representa un libro dentro de la librería virtual.
type Libro struct {
    ID        int
    Titulo    string
    Autor     string
    Disponible bool
}

// Constructor
func NuevoLibro(id int, titulo, autor string) Libro {
    return Libro{
        ID:        id,
        Titulo:    titulo,
        Autor:     autor,
        Disponible: true,
    }
}

// Método
func (l *Libro) MarcarComoPrestado() {
    l.Disponible = false
}

func (l *Libro) MarcarComoDisponible() {
    l.Disponible = true
}

# internal/models/usuario.go
package models

// Usuario representa una persona registrada en la librería.
type Usuario struct {
    ID     int
    Nombre string
}

// Constructor
func NuevoUsuario(id int, nombre string) Usuario {
    return Usuario{
        ID:     id,
        Nombre: nombre,
    }
}

# internal/models/prestamo.go

package models

import "fmt"

// Prestamo une a un usuario y un libro.
type Prestamo struct {
    UsuarioID int
    LibroID   int
}

// Interface Reproducible para cumplir Unidad 3
type Accion interface {
    Ejecutar() string
}

func (p Prestamo) Ejecutar() string {
    return fmt.Sprintf("Usuario %d está solicitando el libro %d", p.UsuarioID, p.LibroID)
}

2. SERVICIOS (Lógica + Slices + Maps + Errores)

# internal/services/libro_service.go
package services

import (
    "errors"
    "libreria-virtual/internal/models"
)

type LibroService struct {
    libros map[int]models.Libro
}

func NuevoLibroService() *LibroService {
    return &LibroService{
        libros: make(map[int]models.Libro),
    }
}

// CRUD: Registrar libro
func (s *LibroService) AgregarLibro(l models.Libro) error {
    if _, existe := s.libros[l.ID]; existe {
        return errors.New("el libro ya existe")
    }
    s.libros[l.ID] = l
    return nil
}

// CRUD: Buscar libro
func (s *LibroService) BuscarLibro(id int) (models.Libro, error) {
    libro, existe := s.libros[id]
    if !existe {
        return models.Libro{}, errors.New("libro no encontrado")
    }
    return libro, nil
}

// CRUD: Listar libros
func (s *LibroService) ListarLibros() []models.Libro {
    lista := []models.Libro{}
    for _, v := range s.libros {
        lista = append(lista, v)
    }
    return lista
}

# internal/services/usuario_service.go
package services

import (
    "errors"
    "libreria-virtual/internal/models"
)

type UsuarioService struct {
    usuarios map[int]models.Usuario
}

func NuevoUsuarioService() *UsuarioService {
    return &UsuarioService{
        usuarios: make(map[int]models.Usuario),
    }
}

func (s *UsuarioService) AgregarUsuario(u models.Usuario) error {
    if _, existe := s.usuarios[u.ID]; existe {
        return errors.New("el usuario ya existe")
    }
    s.usuarios[u.ID] = u
    return nil
}

func (s *UsuarioService) BuscarUsuario(id int) (models.Usuario, error) {
    u, existe := s.usuarios[id]
    if !existe {
        return models.Usuario{}, errors.New("usuario no encontrado")
    }
    return u, nil
}

# internal/services/prestamo_service.go
package services

import (
    "errors"
    "libreria-virtual/internal/models"
)

type PrestamoService struct {
    prestamos []models.Prestamo
    libros    *LibroService
    usuarios  *UsuarioService
}

func NuevoPrestamoService(ls *LibroService, us *UsuarioService) *PrestamoService {
    return &PrestamoService{
        prestamos: []models.Prestamo{},
        libros:    ls,
        usuarios:  us,
    }
}

func (s *PrestamoService) CrearPrestamo(usuarioID, libroID int) error {
    usuario, err := s.usuarios.BuscarUsuario(usuarioID)
    if err != nil {
        return err
    }

    libro, err := s.libros.BuscarLibro(libroID)
    if err != nil {
        return err
    }

    if !libro.Disponible {
        return errors.New("el libro ya está prestado")
    }

    // Actualizar estado del libro
    libro.MarcarComoPrestado()
    s.libros.libros[libroID] = libro

    prestamo := models.Prestamo{
        UsuarioID: usuario.ID,
        LibroID:   libro.ID,
    }

    s.prestamos = append(s.prestamos, prestamo)

    return nil
}

3. ARCHIVO PRINCIPAL (main.go)
package main

import (
    "fmt"
    "libreria-virtual/internal/models"
    "libreria-virtual/internal/services"
)

func main() {

    libroService := services.NuevoLibroService()
    usuarioService := services.NuevoUsuarioService()
    prestamoService := services.NuevoPrestamoService(libroService, usuarioService)

    // Crear usuarios
    usuario := models.NuevoUsuario(1, "Samuel Riera")
    usuarioService.AgregarUsuario(usuario)

    // Crear libros
    libro := models.NuevoLibro(101, "El señor de los anillos", "Tolkien")
    libroService.AgregarLibro(libro)

    // Préstamo
    err := prestamoService.CrearPrestamo(1, 101)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Préstamo realizado con éxito")
    }

    fmt.Println("Libros:", libroService.ListarLibros())
}

🎯 Objetivo del Proyecto

El objetivo de este proyecto es integrar todos los conocimientos adquiridos durante las 8 semanas del curso de Programación con Go, mediante el desarrollo de una Biblioteca Virtual, aplicando conceptos fundamentales del lenguaje Go, estructuras de datos, programación orientada a objetos, concurrencia y servicios web.

El sistema permite simular el funcionamiento de una biblioteca digital, gestionando libros y usuarios a través de una aplicación web conectada a una base de datos MySQL.

📌 Justificación del Tema

El tema Biblioteca Virtual fue seleccionado porque:

Representa un sistema real utilizado en entornos educativos y sociales.

Permite aplicar progresivamente los temas del sílabo.

Integra backend, base de datos y servicios web.

Facilita la visualización del impacto de las tecnologías en la sociedad.

🧠 Integración del Proyecto con el Sílabo

A continuación se detalla cómo cada unidad, tema y semana del curso fue aplicada directamente en el proyecto.

🟦 UNIDAD 1 – Programación con Go
🔹 Semana 1 – TEMA 1: ¿Qué es Go?

Contenidos del sílabo:

Sintaxis

Condicionales

Estructuras de control de flujo iterativo

Aplicación en el proyecto:

Uso de la sintaxis básica de Go en todos los archivos .go.

Uso de condicionales (if, manejo de errores).

Uso de estructuras iterativas como for para recorrer resultados de la base de datos.

Comprensión del flujo de ejecución del servidor web.

Ejemplo aplicado:

Validación de errores al conectar con MySQL.

Recorrido de filas (rows.Next()) en consultas SQL.

🔹 Semana 2 – TEMA 2: Manejo de Funciones y Paquetes

Contenidos del sílabo:

Creación y llamado de funciones

Tipos de funciones

Uso de paquetes

Aplicación en el proyecto:

Creación de funciones como main(), Connect(), Home(), Login(), etc.

Separación del código en paquetes:

db

handlers

models

Uso de import para organizar y reutilizar código.

Modularización del sistema para mejorar mantenimiento y legibilidad.

🟩 UNIDAD 2 – Estructuras de Datos y Objetos
🔹 Semana 3 – TEMA 1: Arrays, Slices y Maps

Contenidos del sílabo:

Manejo de arrays

Manejo de slices

Manejo de maps

Aplicación en el proyecto:

Uso de slices para almacenar listas de libros.

Manejo dinámico de datos obtenidos desde la base de datos.

Uso implícito de estructuras dinámicas para manejar múltiples registros.

Ejemplo:

Slice de libros []Book que se envía como JSON al cliente.

🔹 Semana 4 – TEMA 2: Objetos en Go

Contenidos del sílabo:

Structs

Métodos

Constructores

Aplicación en el proyecto:

Uso de structs para representar entidades del sistema:

Usuario

Libro

Modelado de datos utilizando structs como objetos.

Uso de estructuras para serialización JSON.

Representación clara de los datos del dominio del sistema.

🟨 UNIDAD 3 – Programación Orientada a Objetos
🔹 Semana 5 – TEMA 1: Encapsulación

Contenidos del sílabo:

Métodos setter

Manejo de errores

Aplicación en el proyecto:

Encapsulación de la lógica de conexión a base de datos dentro del paquete db.

Manejo de errores en conexiones, consultas y respuestas HTTP.

Protección de la lógica interna del sistema mediante paquetes internos (internal).

🔹 Semana 6 – TEMA 2: Interfaces

Contenidos del sílabo:

Creación de interfaces

Implementación

Polimorfismo

Aplicación en el proyecto:

Uso indirecto de interfaces propias del lenguaje Go (por ejemplo http.Handler).

Comprensión del polimorfismo aplicado a handlers web.

Preparación del sistema para futuras extensiones mediante interfaces.

🟥 UNIDAD 4 – Concurrencia, Testing y Web
🔹 Semana 7 – TEMA 1: Concurrencia

Contenidos del sílabo:

Introducción a la concurrencia

Goroutines

Canales

Aplicación en el proyecto:

Uso del servidor HTTP de Go que maneja múltiples solicitudes concurrentes.

Comprensión de cómo Go maneja múltiples usuarios al mismo tiempo.

Preparación conceptual para sistemas escalables.

🔹 Semana 8 – TEMA 2: Web

Contenidos del sílabo:

Servicios Web

Serialización de datos

Testing

Aplicación en el proyecto:

Implementación de Servicios Web REST.

Uso de JSON como formato de serialización.

Respuestas HTTP estructuradas.

Consumo de servicios desde el navegador.

Base para futuras pruebas (testing).

🌐 Servicios Web Implementados

El proyecto implementa al menos 8 servicios web, entre ellos:

Servicio Home

Servicio Login

Servicio Registro

Servicio Listado de Libros

Servicio de conexión a base de datos

Servicio de consulta SQL

Servicio de serialización JSON

Servicio de respuesta HTTP

🗂️ Estructura del Proyecto
biblioteca-virtual-go

biblioteca-virtual-go
│── go.mod
│── go.sum
│── README.md
│
├── cmd
│ └── server
│ └── main.go
│
├── internal
│ ├── db
│ │ └── db.go
│ ├── handlers
│ │ ├── home.go
│ │ ├── auth.go
│ │ └── books.go
│ └── models
│ ├── user.go
│ └── book.go
│
├── templates
│ ├── home.html
│ ├── login.html
│ └── register.html
│
└── sql
└── biblioteca.sql

🛢️ Base de Datos

La base de datos MySQL contiene:

Tabla de usuarios

Tabla de libros

Permite almacenar y recuperar información de forma persistente.

🔮 Visualización del Futuro (Unidad 4 – Evaluación Final)

Este proyecto puede evolucionar hacia:

Plataformas educativas digitales

Sistemas bibliotecarios reales

Aplicaciones móviles

Integración con inteligencia artificial

Sistemas en la nube

Microservicios

La Biblioteca Virtual representa el impacto positivo de las nuevas tecnologías en la sociedad, facilitando el acceso al conocimiento.

🧠 Conclusión

El desarrollo de este proyecto permitió aplicar de forma práctica y progresiva todos los temas del sílabo, consolidando conocimientos en Go, programación estructurada, POO, servicios web y bases de datos.

La Biblioteca Virtual demuestra cómo el lenguaje Go puede utilizarse para construir sistemas modernos, eficientes y escalables.

