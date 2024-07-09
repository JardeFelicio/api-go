package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var dbPath = "./sqlite.db"

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	fmt.Printf("Connected to SQLite database at %s\n", dbPath)
	return db, nil
}

func CloseDB(db *sql.DB) {
	if db != nil {
		err := db.Close()
		if err != nil {
			fmt.Println("Error closing database:", err)
		} else {
			fmt.Println("Database connection closed.")
		}
	}
}
