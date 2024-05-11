package main

import (
	"fmt"
	"os"
	"strings"

	ascii "ascii-art/banner"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Errror: Invalid number of argument")
		os.Exit(1)
	}
	input := os.Args[1]

	input = strings.ReplaceAll(input, "\\n", "\n")
	// Removing special cases and returning an error massage.
	input = ascii.RemoveSpecialCases(input)
	if input == "\n" {
		fmt.Println()
	}
	if input == "" {
		return
	}
	// Split the input based on newline
	inputFile := strings.Split(input, "\n")
	var countSpace int
	for _, word := range inputFile {
		if word == "" {
			countSpace++
			if countSpace < len(input) {
				countSpace++
			}
		} else {
			// Print Banner 
			ascii.PrintBanner(word)
		}
	}
}
