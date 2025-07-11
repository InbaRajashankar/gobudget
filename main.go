package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}
	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument")
		return
	}

	input := os.Args[1]
	fmt.Println("You entered: ", input)
}
