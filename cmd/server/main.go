package main

import (
	"log"
	"net/http"

	"biblioteca-virtual/internal/db"
	"biblioteca-virtual/internal/handlers"
)

func main() {

	// Conexión a la base de datos
	db.Connect()

	// Rutas del sistema
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/books", handlers.BooksAPI)

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
