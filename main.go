package main

import (
	"log"

	"github.com/InbaRajashankar/gobudget/backend"
)

func main() {
	args := make(map[string]string)
	args["-n"] = "20"
	args["-p"] = "10,20"
	args["-d"] = "2/12/2023,2/12/2023"
	rows, err := backend.SelectFilter(args)
	if err != nil {
		log.Fatal(err)
	}
	backend.PrintRows(rows)
}
