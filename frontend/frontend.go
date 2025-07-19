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
		case "grab":
			acceptable_args := []string{"-d", "-p", "-n", "-csv"}
			args, err := CreateArgMap(command_arr, acceptable_args)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				HandleGrab(args)
			}
		case "grabsum":
			fmt.Println("grabsum")
		case "graph":
			fmt.Println("graph")
		case "enter":
			fmt.Println("enter")
		default:
			fmt.Println("Invalid command, please enter h for help.")
		}

	}
}

// CreateArgMap creates a map storing [arg]value, only including specified flags
func CreateArgMap(arr []string, acceptable_args []string) (map[string]string, error) {
	args := make(map[string]string)
	for ind, val := range arr {
		if slices.Contains(acceptable_args, val) {
			_, already_in_map := args[val]
			if already_in_map {
				return nil, errors.New("invalid input: duplicate tags")
			}
			if ind+1 >= len(arr) {
				return nil, errors.New("invalid input: no value for final arg")
			}
			args[val] = arr[ind+1]
		}
	}
	return args, nil
}

// HandleGrab is the handler if the command is "grab".
func HandleGrab(args map[string]string) {
	if len(args) == 0 {
		rows := backend.GrabAll()
		backend.PrintRows(rows)
	} else {
		rows, err := backend.GrabFilter(args)
		if err != nil {
			fmt.Println(err)
		} else {
			backend.PrintRows(rows)
		}
	}
}
