package handlers

import (
	"encoding/json"
	"net/http"

	"biblioteca-virtual/internal/db"
	"biblioteca-virtual/internal/models"
)

func BooksAPI(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query("SELECT id, title, author FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	books := []models.Book{}

	for rows.Next() {
		var book models.Book
		rows.Scan(&book.ID, &book.Title, &book.Author)
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
