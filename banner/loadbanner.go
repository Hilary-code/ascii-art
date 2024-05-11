package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadBanner(name string) map[rune]string {
	var height int
	Banner := make(map[rune]string)
	currentChar := rune(32)
	charLine := []string{}
	filepath := "./" + name + ".txt"

	// open file
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error occurred during file opening", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // To skip the first line
	for scanner.Scan() {
		line := scanner.Text()
		if height == 8 {
			Banner[currentChar] = strings.Join(charLine, "\n")
			currentChar++
			height = 0
			charLine = []string{}
		} else {
			charLine = append(charLine, line)
			height++
		}
	}
	if height > 0 {
		Banner[currentChar] = strings.Join(charLine, "\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error occurred during scanning")
		os.Exit(1)
	}
	return Banner
}
