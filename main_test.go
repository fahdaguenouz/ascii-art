package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

func TestMain_NoArgs(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go") // Adjust if your filename is different
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()

	if err == nil {
		t.Errorf("Expected an error when no arguments are passed, but got none")
	}
	expected := "Please enter a string to print."
	if !bytes.Contains(out.Bytes(), []byte(expected)) {
		t.Errorf("Expected output to contain %q, but got:\n%s", expected, out.String())
	}
}

func TestMain_TooManyArgs(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "arg1", "arg2") // Simulate too many arguments
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()

	if err == nil {
		t.Errorf("Expected an error when too many arguments are passed, but got none")
	}
	expected := "Too many arguments. Enter just the string to print."
	if !bytes.Contains(out.Bytes(), []byte(expected)) {
		t.Errorf("Expected output to contain %q, but got:\n%s", expected, out.String())
	}
}

func TestMain_ValidArgs(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "Hello") // Simulate valid input
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()

	if err != nil {
		t.Errorf("Expected no error for valid input, but got: %v", err)
	}
	// Check for non-empty output
	if out.Len() == 0 {
		t.Error("Expected some output, but got none.")
	}
}

func TestMain_FileNotFound(t *testing.T) {
	// Store the original file path
	originalFileName := "./art/standard.txt"

	// Create a temporary file to move the original file to
	tempFile, err := ioutil.TempFile("", "standard.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up the temporary file

	// Move the original file to the temporary location
	err = os.Rename(originalFileName, tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to move the original file: %v", err)
	}
	defer os.Rename(tempFile.Name(), originalFileName) // Restore the original file

	cmd := exec.Command("go", "run", "main.go", "Hello") // Simulate valid input
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err = cmd.Run()

	if err == nil {
		t.Errorf("Expected an error when the file does not exist, but got none")
	}
	expected := "Error opening the file:"
	if !bytes.Contains(out.Bytes(), []byte(expected)) {
		t.Errorf("Expected output to contain %q, but got:\n%s", expected, out.String())
	}
}
