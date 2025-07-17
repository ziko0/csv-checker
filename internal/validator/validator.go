package validator

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ValidateCSV runs all basic validation checks on a CSV file
func ValidateCSV(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading CSV: %w", err)
	}

	if len(rows) == 0 {
		return fmt.Errorf("CSV file is empty")
	}

	expectedCols := len(rows[0])

	for i, row := range rows {
		if err := checkColumnCount(row, expectedCols, i); err != nil {
			return err
		}
		if err := checkEmptyFields(row, i); err != nil {
			return err
		}
		if err := checkForbiddenCharacters(row, i); err != nil {
			return err
		}
	}

	return nil
}
