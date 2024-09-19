package main

import (
	
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		input := args[0]
		fileName:= "./art/standard.txt"
		fmt.Print(input)
		art, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening the file:", err)
			return
		}
		defer art.Close()
		

		
		
	} else if len(args) > 2 {
		fmt.Println("Too much arguments enter just the string to print ")
		os.Exit(1)
	} else if len(args) < 1 {
		fmt.Println("Less arguments please  enter the string to print it ")
		os.Exit(1)
	}
}