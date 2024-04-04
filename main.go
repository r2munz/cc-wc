package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("This is my implementation of wc")
	// Define flags to build -h
	countLines := flag.Bool("l", false, "print the newline counts")
	countWords := flag.Bool("w", false, "print the word counts")
	countChars := flag.Bool("m", false, "print the character counts")
	countBytes := flag.Bool("c", false, "print the bytes counts")
	countMaxLineLength := flag.Bool("L", false, "print the length of the longest line ")
	// Parse command-line arguments
	flag.Parse()
	// step 5 if no flags are provided, default to count lines, words, and characters
	if !(*countLines || *countWords || *countBytes || *countChars || *countMaxLineLength) {
		*countLines, *countWords, *countBytes = true, true, true
	}
	totalLines := 0
	totalWords := 0
	totalChars := 0
	totalBytes := 0
	totalMaxLineLength := 0
	if flag.NArg() == 0 {
		// If no arguments are provided, process stdin
		err := processStdin(countLines, countWords, countBytes, countChars, countMaxLineLength, totalLines, totalWords, totalChars, totalBytes, totalMaxLineLength)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Process each file provided as an argument
		for _, filename := range flag.Args() {
			err := processFile(filename, countLines, countWords, countBytes, countChars, countMaxLineLength, totalLines, totalWords, totalChars, totalBytes, totalMaxLineLength)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error processing file %s: %v\n", filename, err)
				os.Exit(1)
			}
		}
	}
}

func processFile(
	filename string,
	countLines, countWords, countBytes, countChars, countMaxLineLength *bool,
	totalLines, totalWords, totalBytes, totalChars, totalMaxLineLength int,
) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error parsing data: %v\n", err)
	}

	data, err = os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error parsing data: %v\n", err)
	}

	dataString := string(data)
	// step 1 calculate number of bytes
	bytes := len(data)
	// step 3 calculate number of words
	words := CountWords(dataString)
	// step 4 calculate number of characters
	chars := len([]rune(dataString))
	// step 2 lines
	// Initiate NewScanner object with our file
	scanner := bufio.NewScanner(file)
	lines, maxChars := 0, 0
	for scanner.Scan() {
		lines++
		if maxChars < len(scanner.Text()) {
			maxChars = len(scanner.Text())
		}
	}
	// Print counts based on provided flags on the same format as wc does
	if *countLines {
		fmt.Printf("%d\t", lines)
		totalLines += lines
	}
	if *countWords {
		fmt.Printf("%d\t", words)
		totalWords += words
	}
	if *countBytes {
		fmt.Printf("%d\t", bytes)
		totalBytes += int(bytes)
	}
	if *countChars {
		fmt.Printf("%d\t", chars)
		totalChars += chars
	}
	if *countMaxLineLength {
		fmt.Printf("%d\t", maxChars)
		totalMaxLineLength += maxChars
	}
	fmt.Printf("%s\n", filename)

	return nil
}

func processStdin(
	countLines, countWords, countBytes, countChars, countMaxLineLength *bool,
	totalLines, totalWords, totalBytes, totalChars, totalMaxLineLength int,
) error {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	dataString := string(data)
	// step 1 calculate number of bytes
	bytes := len(data)
	// step 3 calculate number of words
	words := CountWords(dataString)
	// step 4 calculate number of characters
	chars := len([]rune(dataString))
	// step 2 lines
	// Initiate NewScanner object with our file
	scanner := bufio.NewScanner(os.Stdin)
	lines, maxChars := 0, 0
	for scanner.Scan() {
		lines++
		if maxChars < len(scanner.Text()) {
			maxChars = len(scanner.Text())
		}
	}
	// Print counts based on provided flags on the same format as wc does
	if *countLines {
		fmt.Printf("%d\t", lines)
		totalLines += lines
	}
	if *countWords {
		fmt.Printf("%d\t", words)
		totalWords += words
	}
	if *countBytes {
		fmt.Printf("%d\t", bytes)
		totalBytes += int(bytes)
	}
	if *countChars {
		fmt.Printf("%d\t", chars)
		totalChars += chars
	}
	if *countMaxLineLength {
		fmt.Printf("%d\t", maxChars)
		totalMaxLineLength += maxChars
	}
	fmt.Printf("\n")
	return nil
}

func CountWords(s string) int {
	return len(strings.Fields(s))
}
