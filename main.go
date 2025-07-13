package main

import "github.com/InbaRajashankar/gobudget/backend"

func main() {
	// args := make(map[string]string)
	// args["-n"] = "2"
	rows := backend.SelectAll()
	// rows := backend.SelectFilter(args)
	backend.PrintRows(rows)

}
