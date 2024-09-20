package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		if len(args) > 1 {
			fmt.Println("Too many arguments. Enter just the string to print.")
		} else {
			fmt.Println("Please enter a string to print.")
		}
		os.Exit(1)
	}

	input := args[0]
	fileName := "./art/standard.txt"

	// Open the ASCII art file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()


	var asciiArt []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		asciiArt = append(asciiArt, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}


	const numLinesPerChar = 8
	startChar := 32
	spaceBetweenChars := 1     // One space between the ASCII characters

	// Each character has an empty line before and after the 8 lines
	linesPerCharacter := numLinesPerChar + spaceBetweenChars // 1 empty line before and after the 8 lines

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		// Loop over the 8 lines for each character row
		for i := 0; i < numLinesPerChar; i++ {
			// Loop over each character in the line
			for _, char := range line {
				asciiIndex := int(char) - startChar // Find the index of the character in the file
				lineIndex := asciiIndex*linesPerCharacter + spaceBetweenChars + i // Account for empty lines

				// Check if the index is valid and print the corresponding line for the character
				if lineIndex < len(asciiArt) {
					fmt.Print(asciiArt[lineIndex])
				}
				fmt.Print(" ") // Space between characters
			}
			fmt.Println() // Newline after each row of characters
		}
		fmt.Println() // Extra newline after processing each line in the input
	}
}