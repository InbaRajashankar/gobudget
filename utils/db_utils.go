package utils

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTable(path string) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `
		CREATE TABLE IF NOT EXISTS lineitems (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			day INTEGER,
			month INTEGER,
			year INTEGER,
			name TEXT,
			price REAL,
			tag TEXT
		);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table created successfully")
}

func AddEntry(day int, month int, year int, name string, price float64, tag string) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec(
		"INSERT INTO lineitems(day, month, year, name, price, tag) VALUES(?, ?, ?, ?, ?, ?)",
		day, month, year, name, price, tag)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("New entry successfully")
}
