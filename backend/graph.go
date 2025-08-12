package backend

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// PlotSqlRows creates a plot from a sql.rows, the output of a Grabsum call.
func PlotSqlRows(rows *sql.Rows, args map[string]string) error {
	vals_map := make(map[string]float64) // map that stores the values
	min_net := 0.0
	max_net := 0.0

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

		var category_name string

		if group_by_tag && group_by_month {
			err = rows.Scan(&count, &net, &tag, &month, &year)
			if err != nil {
				log.Fatal(err)
			}
			category_name = tag + "_" + strconv.Itoa(month) + "_" + strconv.Itoa(year)

		} else if group_by_tag && group_by_year {
			err = rows.Scan(&count, &net, &tag, &year)
			if err != nil {
				log.Fatal(err)
			}
			category_name = tag + "_" + strconv.Itoa(year)

		} else if group_by_tag {
			err = rows.Scan(&count, &net, &tag)
			if err != nil {
				log.Fatal(err)
			}
			category_name = tag

		} else if group_by_month {
			err = rows.Scan(&count, &net, &month, &year)
			if err != nil {
				log.Fatal(err)
			}
			category_name = strconv.Itoa(month) + "_" + strconv.Itoa(year)

		} else if group_by_year {
			err = rows.Scan(&count, &net, &year)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%d %.2f %d", count, net, year)
			category_name = strconv.Itoa(year)

		} else {
			err = rows.Scan(&count, &net)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%d %.2f", count, net)
			category_name = ""
		}

		min_net = min(min_net, net)
		max_net = max(max_net, net)
		vals_map[category_name] = net

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	GeneratePlot(&vals_map, min_net, max_net)

	return nil
}

func GeneratePlot(vals_map *map[string]float64, min_net float64, max_net float64) error {
	range_of_xs := 50.0 // width of xs on the screen
	net_range := max_net - min_net
	below_zero_range_xs := 0
	if min_net < 0 {
		below_zero_range_xs = int(math.Round(-1 * min_net * range_of_xs / net_range))
	}
	var count_xs int
	for key, val := range *vals_map {
		if val < 0 {
			count_xs = int(math.Round((-1 * val / net_range) * range_of_xs))
			fmt.Printf("|%15s|%10.2f|", key, val)
			fmt.Print(strings.Repeat(" ", below_zero_range_xs-count_xs+1))
			fmt.Println("\033[35m" + strings.Repeat("x", count_xs) + "\033[0m|")
		} else {
			count_xs = int(math.Round((val / net_range) * range_of_xs))
			fmt.Printf("|%15s|%10.2f|", key, val)
			fmt.Print(strings.Repeat(" ", below_zero_range_xs+1) + "|")
			fmt.Println("\033[36m" + strings.Repeat("x", count_xs) + "\033[0m")
		}
	}

	return nil
}
