package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL") // contoh: postgres://user:password@localhost:5432/golang_book?sslmode=disable
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connect database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error ping database:", err)
	}

	fmt.Println("Database connected ✅")
}
