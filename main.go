package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	// validate the input parameters
	if len(os.Args) != 2 {
		fmt.Println("Please specifiy a file name")
		os.Exit(2)
	}

	// get the filepath parameter
	filepath := os.Args[1]
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("File %q not found.\n", filepath)
		os.Exit(2)
	}

	// split into lines
	lines := getLines(file)

	for _, line := range lines {

		normalizedLine := []byte(line)

		// replace the session id
		sessionIdPattern := regexp.MustCompile(`<sessionId>[^<]+?</sessionId>`)
		normalizedLine = sessionIdPattern.ReplaceAll(normalizedLine, []byte(`<sessionId>?</sessionId>`))

		// remove white space
		normalizedLine = bytes.TrimSpace(normalizedLine)

		// print the normalized line
		if len(normalizedLine) > 0 {
			fmt.Printf("%s\n", normalizedLine)
		}
	}
}

// readLine returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func readLine(bufferedReader *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)

	for isPrefix && err == nil {
		line, isPrefix, err = bufferedReader.ReadLine()
		ln = append(ln, line...)
	}

	return string(ln), err
}

// Get all lines of a given file
func getLines(inFile io.Reader) []string {

	lines := make([]string, 0, 10)
	bufferedReader := bufio.NewReader(inFile)
	line, err := readLine(bufferedReader)
	for err == nil {
		lines = append(lines, line)
		line, err = readLine(bufferedReader)
	}

	return lines
}
