package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	var pathCount int
	var filePath string
	var keepQuery bool

	flag.IntVar(&pathCount, "p", 0, "Number of path segments to keep in the URL")
	flag.StringVar(&filePath, "f", "", "File path for reading URLs")
	flag.BoolVar(&keepQuery, "keep-query", false, "Keep the query parameters in the URL")
	flag.Parse()

	var scanner *bufio.Scanner
	if filePath != "" {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	for scanner.Scan() {
		processedURL, err := processURL(scanner.Text(), pathCount, keepQuery)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error processing URL: %v\n", err)
			continue
		}
		fmt.Println(processedURL)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}

func processURL(inputURL string, pathCount int, keepQuery bool) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	if pathCount >= 0 {
		segments := strings.Split(parsedURL.Path, "/")
		if len(segments) > pathCount {
			parsedURL.Path = strings.Join(segments[:pathCount+1], "/")
		}
	}

	if !keepQuery {
		parsedURL.RawQuery = ""
	}

	return parsedURL.String(), nil
}
