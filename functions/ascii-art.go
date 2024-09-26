package asciiart

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PrintAsciiArt(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("Please enter a string to print.")
	}

	fileName := "./art/standard.txt"

	// Open the ASCII art file
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("Error opening the file: %v", err)
	}
	defer file.Close()

	var asciiArt []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		asciiArt = append(asciiArt, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("Error reading the file: %v", err)
	}

	var result string
	lines := strings.Split(input, "\\n")

	for _, line := range lines {
		if line == "" {
			result += "\n"
			continue
		}
		// Iterate over each row of the ASCII art (0 to 7, for the 8 rows)
		for i := 1; i <= 8; i++ {
			for _, r := range line {
				// Ensure the character is within the valid ASCII range
				if r < 32 || r > 126 {
					return "", fmt.Errorf("Please enter a valid character between ascii code 32 and 126")
				}
				index := 9*(int(r)-32) + i
				result += asciiArt[index]
			}
			result += "\n" // Add newline after finishing the current row of the line
		}
	}
	return result, nil
}
