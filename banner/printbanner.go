package ascii

import (
	"fmt"
	"os"
	"strings"
)

func PrintBanner(line string) {
	output := make([][]string, 8)
	var fileName string

	if len(os.Args) == 3 {
		fileName = os.Args[2]
	} else {
		fileName = "standard" // default file name if the file name is not spacified
	}
	filepath := fileName

	// Check if the size of banner has been altered.
	filepath, err := FileCheck(fileName)
	if err != nil {
		fmt.Println(err, fileName)
		os.Exit(1)
	}

	banner := LoadBanner(filepath)
	for _, char := range line {
		if char < 32 && char > 126 {
			fmt.Printf("Character out of range: %q\n", char)
			os.Exit(1)
		}
		if asciiInput, Ok := banner[char]; Ok {
			asciiLine := strings.Split(asciiInput, "\n")
			for i := 0; i < len(asciiLine); i++ {
				output[i] = append(output[i], asciiLine[i])
			}
		} else {
			fmt.Printf("Character not found: %q\n", char)
			continue
		}
	}
	for _, line := range output {
		fmt.Println(strings.Join(line, ""))
	}
}
