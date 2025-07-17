package validator

import (
	"fmt"
	"strings"
)

// checkColumnCount ensures each row has the same number of columns as the header
func checkColumnCount(row []string, expected int, rowIndex int) error {
	if len(row) != expected {
		return fmt.Errorf("row %d: expected %d columns, found %d", rowIndex+1, expected, len(row))
	}
	return nil
}

// checkEmptyFields flags any empty fields in a row
func checkEmptyFields(row []string, rowIndex int) error {
	for i, val := range row {
		if strings.TrimSpace(val) == "" {
			return fmt.Errorf("row %d: column %d is empty", rowIndex+1, i+1)
		}
	}
	return nil
}

// checkForbiddenCharacters detects forbidden special characters in a field
// NOTE: '@' is allowed and NOT considered a forbidden character
func checkForbiddenCharacters(row []string, rowIndex int) error {
	// Define forbidden characters (excluding @)
	const forbidden = `#$%^&*()+=[]{}:;"<>\/|~`

	for i, val := range row {
		if strings.ContainsAny(val, forbidden) {
			return fmt.Errorf("row %d: column %d contains forbidden characters", rowIndex+1, i+1)
		}
	}
	return nil
}
