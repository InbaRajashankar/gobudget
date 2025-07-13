package backend

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/InbaRajashankar/gobudget/utils"

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
func SelectFilter(args map[string]string) (*sql.Rows, error) {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "SELECT * FROM expenses"
	conditions_applied := 0 // in case an AND is needed in between multiple BETWEEN statements

	// filter by time
	date_range, ok := args["-d"]
	if ok {
		suffix, err := DateRangeToQueryString(date_range)
		if err != nil {
			return nil, err
		}
		query += suffix
		conditions_applied++
	}

	// filter by price
	price_range, ok := args["-p"]
	price_arr := strings.Split(price_range, ",")
	if ok {
		if len(price_arr) != 2 || price_arr[0] > price_arr[1] {
			return nil, errors.New("invalid price range")
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
		query += " LIMIT " + n_rows
	}

	log.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	return rows, nil
}

// M/D/Y,M/D/Y ->
func DateRangeToQueryString(range_str string) (string, error) {
	// Process string into dates, validation
	times := strings.Split(range_str, ",")
	if len(times) != 2 {
		return "", errors.New("< 2 dates provided")
	}

	d1, err := utils.StringToDateValues(times[0])
	if err != nil {
		return "", errors.New("invalid first date")
	}

	d2, err := utils.StringToDateValues(times[1])
	if err != nil {
		return "", errors.New("invalid second date")
	}

	// Check if d1 < d2
	valid_range := false
	if d1[2] == d2[2] && d1[0] == d2[2] && d1[1] <= d2[1] {
		valid_range = true
	} else if d1[2] == d2[2] && d1[0] <= d2[0] {
		valid_range = true
	} else if d1[1] <= d2[1] {
		valid_range = true
	}
	if !valid_range {
		return "", errors.New("second date in range is before first")
	}

	// form query
	query := " WHERE "
	if d1 == d2 {
		query += "year = " + strconv.Itoa(d1[2])
		query += " AND month = " + strconv.Itoa(d1[0])
		query += " AND day = " + strconv.Itoa(d1[1])
	} else if d1[2] == d2[2] && d1[0] == d2[0] {
		query += "year = " + strconv.Itoa(d1[2])
		query += " AND month = " + strconv.Itoa(d1[0])
		query += " AND day BETWEEN " + strconv.Itoa(d1[1]) + " AND " + strconv.Itoa(d2[1])
	} else if d1[2] == d2[2] {
		query += "year = " + strconv.Itoa(d1[2])
		// tail end of first month
		query += " AND ((month = " + strconv.Itoa(d1[0]) + " AND day >= " + strconv.Itoa(d1[1]) + ") "
		if d2[0]-d1[0] > 1 { // middle months, if months are not one apart
			query += "OR (month BETWEEN " + strconv.Itoa(d1[0]+1) + " AND " + strconv.Itoa(d2[0]-1) + ") "
		}
		// front end of last month
		query += "OR (month = " + strconv.Itoa(d2[0]) + " AND day <= " + strconv.Itoa(d2[1]) + "))"

	} else {
		// tail end of first year
		query += "(year = " + strconv.Itoa(d1[2])
		query += " AND ((month = " + strconv.Itoa(d1[0]) + " AND day >= " + strconv.Itoa(d1[1]) + ") "
		query += "OR (month > " + strconv.Itoa(d1[0]) + ")))"

		// middle years, if years are not one apart
		if d2[2]-d1[2] > 1 {
			query += " OR (year BETWEEN " + strconv.Itoa(d1[2]+1) + " AND " + strconv.Itoa(d2[2]-1) + ")"
		}

		// front end of last year
		query += " OR (year = " + strconv.Itoa(d2[2])
		query += " AND ((month < " + strconv.Itoa(d2[0]) + ")"
		query += " OR (month = " + strconv.Itoa(d2[0]) + " AND day <= " + strconv.Itoa(d2[1]) + ")))"
	}

	return query, nil
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
