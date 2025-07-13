package backend

import (
	"errors"
	"strconv"
	"strings"

	"github.com/InbaRajashankar/gobudget/utils"
)

// Changes a date range in the format M/D/Y,M/D/Y to an SQLite WHERE statement
func DateRangeToQuerySuffix(range_str string) (string, error) {
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
