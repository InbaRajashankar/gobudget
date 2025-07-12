package main

import (
	"fmt"

	"github.com/InbaRajashankar/gobudget/utils"
)

func main() {
	// args := make(map[string]string)
	// args["-n"] = "2"
	// // rows := backend.SelectAll()
	// rows := backend.SelectFilter(args)
	// backend.PrintRows(rows)

	var test_dates = [6]string{"2/2", "0/23/4", "1/244/4", "3/4/51", "2/-2/2314", "-2/4/2"}

	for i, _ := range test_dates {
		_, err := utils.StringToDateValues(test_dates[i])
		if err != nil {
			fmt.Println(err)
		}
	}
}
