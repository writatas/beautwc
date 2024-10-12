package internal

import (
	"fmt"
	"os"
)

// validate a slice of given file paths, if one produces an error the program should return an error
func ValidateFiles(files ...string) []error {
	errors := []error{}
	for _, f := range files {
		_, err := os.ReadFile(f)
		if err != nil {
			file_err := fmt.Errorf("failed to read file: %s", f)
			errors = append(errors, file_err)
		}
	}
	return errors
}
