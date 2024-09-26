package main

import (
	asciiart "asciiart/functions"
	"strings"
	"testing"
)

func TestPrintAsciiArt(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  string
		expectErr bool
	}{
		{
			name:  "valid input",
			input: "Hello",
			expected: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
`, 
			expectErr: false,
		},
		
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := asciiart.PrintAsciiArt(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("PrintAsciiArt() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			// Trim the result for comparison
			result = strings.TrimSpace(result)
			expected := strings.TrimSpace(tt.expected)

			if !tt.expectErr && result != expected {
				t.Errorf("PrintAsciiArt() = %v, want %v", result, expected)
			}
		})
	}
}