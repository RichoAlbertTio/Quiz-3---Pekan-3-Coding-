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
	// Cek apakah DATABASE_URL tersedia
	databaseURL := os.Getenv("DATABASE_URL")
	var dsn string
	if databaseURL != "" {
		dsn = databaseURL
	} else {
		// Gunakan variabel environment Railway/postgres
		dbHost := os.Getenv("PGHOST")
		dbPort := os.Getenv("PGPORT")
		dbUser := os.Getenv("PGUSER")
		dbPassword := os.Getenv("PGPASSWORD")
		dbName := os.Getenv("PGDATABASE")
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	}

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
