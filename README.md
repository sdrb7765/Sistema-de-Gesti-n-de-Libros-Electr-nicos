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

