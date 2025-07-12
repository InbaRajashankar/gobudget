package backend

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const db_path string = "./test.db"

// SelectAll selects all entries from the db.
func SelectAll() *sql.Rows {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM expenses")
	if err != nil {
		log.Fatal(err)
	}

	return rows
}

// SelectFilter selects all items from the db with some filter.
func SelectFilter(args map[string]string) *sql.Rows {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "SELECT * FROM expenses"

	// need error handling here
	n, ok := args["-n"]
	if ok {
		query += " LIMIT " + n
	}

	log.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func PrintRows(rows *sql.Rows) {
	for rows.Next() {
		var id int
		var date string
		var name string
		var price float64
		var tag string
		err := rows.Scan(&id, &date, &name, &price, &tag)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%d %s %s %.2f %s", id, date, name, price, tag)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
