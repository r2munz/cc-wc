package main

import (
	"bufio"
	"flag"
	"fmt"
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
	// Parse command-line arguments
	flag.Parse()
	// step 5 if no flags are provided, default to count lines, words, and characters
	if !(*countLines || *countWords || *countBytes || *countChars) {
		*countLines, *countWords, *countBytes = true, true, true
	}
	totalLines := 0
	totalWords := 0
	totalChars := 0
	totalBytes := 0
	// Process each file provided in the arguments
	for _, filename := range flag.Args() {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", filename, err)
			continue
		}
		defer file.Close()

		//reader := bufio.NewReader(os.Stdin)
		//input, _ := reader.ReadString('\n')
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error parsing data: %v\n", err)
			continue
		}
		/*
			var dataString string
			var bytes int
			if input == "" {
				dataString = string(data)
				bytes = len(data)
				continue
			} else {
				dataString = input
				bytes = len(input)
			}
		*/
		dataString := string(data)
		// step 1 calculate number of bytes
		bytes := len(data)
		// step 3 calculate number of words
		words := CountWords(dataString)
		// step 2 and 4 calculate number of lines and character
		// Initiate NewScanner object with our file
		scanner := bufio.NewScanner(file)
		lines, chars := 0, 0
		for scanner.Scan() {
			lines++
			chars += len(scanner.Text())
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
		fmt.Printf("%s\n", filename)
	}
	// bonus step calculate totals
	if *countLines {
		fmt.Printf("%d\t", totalLines)
	}
	if *countWords {
		fmt.Printf("%d\t", totalWords)
	}
	if *countBytes {
		fmt.Printf("%d\t", totalBytes)
	}
	if *countChars {
		fmt.Printf("%d\t", totalChars)
	}
	fmt.Printf("%s\n", "total")
}

func CountWords(s string) int {
	return len(strings.Fields(s))
}
