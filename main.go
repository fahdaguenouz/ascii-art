package main

import (
	
	"fmt"
	"os"
	"asciiart/functions"
	
)


func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		if len(args) > 1 {
			fmt.Println("Too many arguments. Enter just the string to print.")
		} else {
			fmt.Println("Please enter a string to print.")
		}
		return
	}

	input := args[0]
	result, err := asciiart.PrintAsciiArt(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(result)
}