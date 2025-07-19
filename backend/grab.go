package backend

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const db_path string = "./test.db"

// GrabAll selects all entries from the db.
func GrabAll() *sql.Rows {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM lineitems")
	if err != nil {
		log.Fatal(err)
	}

	return rows
}

// GrabFilter selects all items from the db with some filter.
func GrabFilter(args map[string]string) (*sql.Rows, error) {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "SELECT * FROM lineitems"
	conditions_applied := 0 // in case an AND is needed in between multiple BETWEEN statements

	// filter by time
	date_range, ok := args["-d"]
	if ok {
		suffix, err := DateRangeToQuerySuffix(date_range)
		if err != nil {
			return nil, err
		}
		query += suffix
		conditions_applied++
	}

	// filter by price
	price_range, ok := args["-p"]
	price_arr := strings.Split(price_range, ",")
	first_price, err := strconv.Atoi(price_arr[0])
	if err != nil {
		return nil, errors.New("invalid first price, cannot convert to integer")
	}
	second_price, err := strconv.Atoi(price_arr[1])
	if err != nil {
		return nil, errors.New("invalid second price, cannot convert to integer")
	}
	if ok {
		if len(price_arr) != 2 || first_price > second_price {
			error_string := fmt.Sprintf("invalid price range %s %s", price_arr[0], price_arr[1])
			return nil, errors.New(error_string)
		}
		if conditions_applied > 0 {
			query += " AND price BETWEEN " + price_arr[0] + " AND " + price_arr[1]
		} else {
			query += " WHERE price BETWEEN " + price_arr[0] + " AND " + price_arr[1]
		}
	}

	// top n rows
	n_rows, ok := args["-n"]
	if ok {
		_, err = strconv.Atoi(n_rows) // check if n_rows is an integer
		if err != nil {
			return nil, errors.New("invalid number of entries")
		}
		query += " LIMIT " + n_rows
	}

	log.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	return rows, nil
}

func PrintRows(rows *sql.Rows) {
	for rows.Next() {
		var id int
		var day int
		var month int
		var year int
		var name string
		var price float64
		var tag string
		err := rows.Scan(&id, &day, &month, &year, &name, &price, &tag)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%d %d %d %d %s %.2f %s", id, day, month, year, name, price, tag)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
