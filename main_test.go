func GenerateASCIIArt(text string, banner []string) string {
	var result string
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if line == "" {
			result += "\n"
			continue
		}
		// Iterate over each row of the ASCII art (0 to 7, for the 8 rows)
		for i := 0; i < 8; i++ {
			for _, r := range line {
				// Ensure the character is within the valid ASCII range
				if r < 32 || r > 126 {
					// Or any other placeholder
					continue
				}
				index := 9*(int(r)-32) + i + 1
				result += banner[index]
			}
			result += "\n" // Add newline after finishing the current row of the line
		}
	}
	return result
}
