package common

import (
	"io/ioutil"
	"log"
	"strings"
)

// ReadAndSanitise Read the input file, trimming whitespace from each row and
// ignoring any blank lines
func ReadAndSanitise(filePath string) []string {
	fileRows := ReadFile(filePath)
	return sanitiseRows(fileRows)
}

// ReadFile Read a text file to an array containing each row
func ReadFile(filePath string) []string {
	fileBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input file")
	}
	fileContents := string(fileBytes)
	return strings.Split(fileContents, "\n")
}

func sanitiseRows(fileRows []string) []string {
	sanitisedRows := make([]string, 0, len(fileRows))
	for _, row := range fileRows {
		sanitisedRow := strings.TrimSpace(row)
		if len(sanitisedRow) > 0 {
			sanitisedRows = append(sanitisedRows, sanitisedRow)
		}
	}
	return sanitisedRows
}

// SolutionNotFoundError Error to return when a solution isn't found
type SolutionNotFoundError string

func (e SolutionNotFoundError) Error() string {
	return string(e)
}
