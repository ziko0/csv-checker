package main

import (
	"fmt"
	"os"

	"csv-checker/internal/validator"
)

// Entry point of the application
// Usage: go run main.go <file.csv>
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: csv-checker <file.csv>")
		return
	}

	file := os.Args[1]
	err := validator.ValidateCSV(file)
	if err != nil {
		fmt.Println("Validation failed:", err)
	} else {
		fmt.Println("CSV is valid.")
	}
}
