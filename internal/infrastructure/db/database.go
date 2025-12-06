package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// membuka koneksi PostgreSQL menggunakan godotenv.
func ConnectPostgres() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("[WARN] .env file not found or failed to load")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open connection error: %w", err)
	}

	// Optional config
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(120 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping error: %w", err)
	}

	log.Println("Database Connected")

	return db, nil
}
