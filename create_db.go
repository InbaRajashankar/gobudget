// package main

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/mattn/go-sqlite3"
// )

// func main() {
// 	add_entry("2.3.1232", "Toad", 4.23, "Animal")
// 	add_entry("2.4.1252", "Cat", 4.23, "Animal")
// 	select_all()
// }

// func create_table() {
// 	db, err := sql.Open("sqlite3", "./test.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	sqlStmt := `
// 		CREATE TABLE IF NOT EXISTS expenses (
// 			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
// 			date TEXT,
// 			name TEXT,
// 			price REAL,
// 			tag TEXT
// 		);
// 	`
// 	_, err = db.Exec(sqlStmt)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("Table created successfully")
// }

// func add_entry(date string, name string, price float64, tag string) {
// 	db, err := sql.Open("sqlite3", "./test.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	_, err = db.Exec("INSERT INTO expenses(date, name, price, tag) VALUES(?, ?, ?, ?)", date, name, price, tag)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("New entry successfully")
// }

// func select_all() {
// 	db, err := sql.Open("sqlite3", "./test.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	rows, err := db.Query("SELECT * FROM expenses")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for rows.Next() {
// 		var id int
// 		var date string
// 		var name string
// 		var price float64
// 		var tag string
// 		err = rows.Scan(&id, &date, &name, &price, &tag)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Printf("%d %s %s %.2f %s", id, date, name, price, tag)
// 	}
// 	if err = rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }
