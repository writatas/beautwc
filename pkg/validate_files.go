package pkg

import (
	"fmt"
	"os"
)

// validate a slice of given file paths, if one produces an error the program should return an error
func ValidateFiles(files ...string) []error {
	errors := []error{}
	for _, f := range files {
		_, err := os.Stat(f)
		if err != nil {
			file_err := fmt.Errorf("failed to read file: %s, %v", f, err)
			errors = append(errors, file_err)
		}
	}
	return errors
}
