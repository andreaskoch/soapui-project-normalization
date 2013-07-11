package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

	// read the file
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error while reading file %q.\n", filepath)
		os.Exit(2)
	}

	// a new line seperator which will later be removed
	newlineString := `-fsdjfhs-NEWLINE-dasjkdjasdkl-`

	// replace windows line endings
	newLinePatternWindows := regexp.MustCompile(`\r\n`)
	content = newLinePatternWindows.ReplaceAll(content, []byte(newlineString))

	// replace unix line endings
	newLinePatternUnix := regexp.MustCompile(`\n`)
	content = newLinePatternUnix.ReplaceAll(content, []byte(newlineString))

	// line breaks after tags
	openingTagPattern := regexp.MustCompile(`(<\w+[^>]*?>)`)
	content = openingTagPattern.ReplaceAll(content, []byte(newlineString+`$1`+newlineString))

	// line breaks before closing tags
	closingTagPattern := regexp.MustCompile(`(</[^>]+>)`)
	content = closingTagPattern.ReplaceAll(content, []byte(newlineString+`$1`+newlineString))

	// split into lines
	lines := bytes.Split(content, []byte(newlineString))
	for _, line := range lines {

		normalizedLine := line

		// replace the session id
		sessionIdPattern := regexp.MustCompile(`<sessionId>[^<]+?</sessionId`)
		normalizedLine = sessionIdPattern.ReplaceAll(normalizedLine, []byte(`<sessionId>?</sessionId>`))

		// trim the line
		normalizedLine = bytes.TrimSpace(normalizedLine)

		// print the normalized line
		if len(normalizedLine) > 0 {
			fmt.Printf("%s\n", line)
		}
	}
}
