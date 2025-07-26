package main

import (
	"log"

	"github.com/InbaRajashankar/gobudget/backend"
	"github.com/InbaRajashankar/gobudget/frontend"
)

func main() {
	// DemoGrabsum()
	// DemoGrabsum()
	DemoFrontend()
}

func DemoFrontend() {
	frontend.InteractionLoop()
}

func DemoGrab() {
	args := make(map[string]string)
	args["-n"] = "20"
	args["-p"] = "10,20"
	args["-d"] = "2/12/2023,2/12/2023"
	rows, err := backend.GrabFilter("./test.db", args)
	if err != nil {
		log.Fatal(err)
	}
	backend.PrintRows(rows)
}

func DemoGrabsum() {
	args := make(map[string]string)
	args["RANGE"] = "1/1/2022,5/1/2023"
	// args["-e"] = ""
	// args["-t"] = ""
	args["-m"] = ""
	rows, err := backend.Grabsum("./test.db", args)
	if err != nil {
		log.Fatal(err)
	}
	backend.PrintGrabsumRows(rows, args)
}
