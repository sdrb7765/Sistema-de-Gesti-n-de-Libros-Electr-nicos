package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error

	DB, err = sql.Open("mysql", "root:1753091667(127.0.0.1:3306)/biblioteca_virtual")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error conectando a MySQL")
	}

	log.Println("Conectado a MySQL correctamente")
}
