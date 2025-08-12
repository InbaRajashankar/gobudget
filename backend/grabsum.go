package backend

import (
	"database/sql"
	"errors"
	"log"
)

func Grabsum(db_path string, args map[string]string) (*sql.Rows, error) {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Figure out which columns to group by based on args
	group_columns := ""
	_, group_by_month := args["-m"]
	_, group_by_year := args["-y"]
	_, group_by_tag := args["-t"]

	if group_by_tag && group_by_month {
		group_columns += "tag, month, year"
	} else if group_by_month {
		group_columns += "month, year"
	} else if group_by_tag && group_by_year {
		group_columns += "tag, year"
	} else if group_by_year {
		group_columns += "year"
	} else if group_by_tag {
		group_columns += "tag"
	}

	// Form Query
	var query string
	if group_columns == "" {
		query = "SELECT COUNT(id) AS count, SUM(price) AS net FROM lineitems"
	} else {
		query = "SELECT COUNT(id) AS count, SUM(price) AS net, " + group_columns + " FROM lineitems"
	}

	date_range_string, has_date_range := args["RANGE"]
	if !has_date_range {
		return nil, errors.New("date range missing")
	}
	date_between_suffix, err := DateRangeToQuerySuffix(date_range_string)
	if err != nil {
		return nil, err
	}
	query += date_between_suffix

	// If income-only or expenses-only is specified
	_, income_only := args["-i"]
	_, expenses_only := args["-e"]
	if date_between_suffix == "" && (income_only || expenses_only) {
		query += " WHERE"
	} else if income_only || expenses_only {
		query += " AND"
	}
	if income_only {
		query += " (price > 0)"
	} else if expenses_only {
		query += " (price < 0)"
	}

	// Groupby
	if group_columns != "" {
		query += " GROUP BY " + group_columns
	}

	query += " ORDER BY net"

	log.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	return rows, nil
}

func PrintGrabsumRows(rows *sql.Rows, args map[string]string) {
	_, group_by_tag := args["-t"]
	_, group_by_month := args["-m"]
	_, group_by_year := args["-y"]
	for rows.Next() {
		var count int
		var net float64
		var err error
		var tag string
		var month int
		var year int

		if group_by_tag && group_by_month {
			err = rows.Scan(&count, &net, &tag, &month, &year)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%d %.2f %s %d %d", count, net, tag, month, year)

		} else if group_by_tag && group_by_year {
			err = rows.Scan(&count, &net, &tag, &year)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%d %.2f %s %d", count, net, tag, year)

		} else if group_by_tag {
			err = rows.Scan(&count, &net, &tag)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%d %.2f %s", count, net, tag)

		} else if group_by_month {
			err = rows.Scan(&count, &net, &month, &year)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%d %.2f %d %d", count, net, month, year)

		} else if group_by_year {
			err = rows.Scan(&count, &net, &year)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%d %.2f %d", count, net, year)

		} else {
			err = rows.Scan(&count, &net)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%d %.2f", count, net)

		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
