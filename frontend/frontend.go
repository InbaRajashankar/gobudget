package frontend

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
			fmt.Println("Goodbye :3")
			os.Exit(0)
		case "help", "h":
			log.Fatal("help has not been implemented yet :p")
		case "grab":
			fmt.Println("grab")
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
