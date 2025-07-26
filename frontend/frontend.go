package frontend

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/InbaRajashankar/gobudget/backend"
)

func InteractionLoop() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to gobudget! Enter h for help.")

	for {
		fmt.Print("> ")
		buffer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		buffer = strings.ReplaceAll(buffer, "\n", "")
		command_arr := strings.Split(buffer, " ")

		switch command_arr[0] {
		case "quit", "q":
			fmt.Println("Goodbye...")
			os.Exit(0)
		case "help", "h":
			log.Fatal("help has not been implemented yet :p")
		case "clear", "c":
			log.Fatal("clear has not been implemented yet :p")
		case "grab", "g":
			err := HandleGrab(command_arr)
			if err != nil {
				log.Println(err)
			}
		case "grabsum", "gs":
			err := HandleGrabsum(command_arr)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "graph":
			fmt.Println("graph")
		case "enter":
			fmt.Println("enter")
		default:
			fmt.Println("Invalid command, please enter h for help.")
		}

	}
}

// HandleGrab is the handler if the command is "grab".
func HandleGrab(arr []string) error {
	// Create the map of arguments, only including args in acceptable_args
	args := make(map[string]string)
	acceptable_args := []string{"-d", "-p", "-n", "-csv"}
	for ind, val := range arr {
		if slices.Contains(acceptable_args, val) {
			_, already_in_map := args[val]
			if already_in_map {
				return errors.New("invalid input: duplicate tags")
			}
			if ind+1 >= len(arr) {
				return errors.New("invalid input: no value for final arg")
			}
			args[val] = arr[ind+1]
		}
	}

	// Call the backend.
	if len(args) == 0 {
		rows := backend.GrabAll()
		backend.PrintRows(rows)
	} else {
		rows, err := backend.GrabFilter(args)
		if err != nil {
			return err
		} else {
			backend.PrintRows(rows)
		}
	}

	return nil
}

// HandleGrabsum is the handler if the command is "grabsum"
func HandleGrabsum(arr []string) error {
	// Create the map of arguments, only including args in acceptable_args
	args := make(map[string]string)
	acceptable_args := []string{"-i", "-e", "-t", "-m", "-y"}

	if len(arr) < 2 {
		return errors.New("no date range provided")
	}
	args["RANGE"] = arr[1]

	for _, val := range arr {
		if slices.Contains(acceptable_args, val) {
			_, already_in_map := args[val]
			if already_in_map {
				return errors.New("invalid input: duplicate tags")
			}
			args[val] = ""
		}
	}

	// Call backend.
	rows, err := backend.Grabsum(args)
	if err != nil {
		return err
	} else {
		backend.PrintGrabsumRows(rows, args)
	}

	return nil
}
