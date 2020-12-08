package common

import (
	"io/ioutil"
	"log"
	"strings"
)

// ReadAndSanitiseRows Read the input file, trimming whitespace from each row and
// ignoring any blank lines
func ReadAndSanitiseRows(filePath string) []string {
	fileRows := ReadFileRows(filePath)
	return sanitiseRows(fileRows)
}

// ReadFileRows Read a text file to an array containing each row
func ReadFileRows(filePath string) []string {
	fileContents := ReadFileContents(filePath)
	return strings.Split(fileContents, "\n")
}

// ReadFileContents Read the contents of a file as a single string
func ReadFileContents(filePath string) string {
	fileBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input file")
	}
	return string(fileBytes)
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
