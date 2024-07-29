package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func RunDbServer() {
	var err error
	dbConn := "user=postgres password=31415926 dbname=Test sslmode=disable"
	db, err = sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatalf("ERROR CONNECT TO DB: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("ERROR PINGING DB: %v", err)
	}
	log.Println("Successfully connected to the database")
}

// AddMessageToDB inserts a message into the test3 table
func AddMessageToDB(message string) bool {
	query := `INSERT INTO test3 (message) VALUES ($1)`
	_, err := db.Exec(query, message)
	if err != nil {
		log.Printf("ERROR INSERTING DATA: %v", err)
		return false
	}
	return true
}
