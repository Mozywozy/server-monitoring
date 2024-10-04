package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

// InitDB menginisialisasi koneksi ke database PostgreSQL
func InitDB() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Pastikan database dapat terhubung
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Membuat tabel jika belum ada
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS server_status (
		id SERIAL PRIMARY KEY,
		server_name VARCHAR(255) NOT NULL,
		status VARCHAR(10) NOT NULL,
		checked_at TIMESTAMP NOT NULL DEFAULT NOW()
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

// SaveServerStatus menyimpan status server ke database
func SaveServerStatus(serverName string, status string) {
	_, err := db.Exec("INSERT INTO server_status (server_name, status, checked_at) VALUES ($1, $2, NOW())", serverName, status)
	if err != nil {
		fmt.Println("Error saving server status:", err)
	}
}

func GetAllServerStatus() (*sql.Rows, error) {
	rows, err := db.Query("SELECT server_name, status FROM server_status ORDER BY checked_at DESC")
	if err != nil {
		return nil, err
	}
	return rows, nil
}