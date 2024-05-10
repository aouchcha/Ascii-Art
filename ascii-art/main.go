package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"ascii-art/tools"
)

func DrawAsciiArt(elements []string) string {
	var result string
	words := strings.Split(os.Args[1], "\\n")
	for _, word := range words {
		// replace empty string by new line
		if word != "" {
			for j := 0; j < 8; j++ {
				for _, char := range word {
					if char < 32 || char > 126 {
						log.Fatalln("Error: please provide printable characters!!\nhelp: man ascii")
					} else {
						// detect the line from where we should start reading
						start := int(char-32)*8 + j
						result += (elements[start])
					}
				}
				result += "\n"
			}
		} else {
			result += "\n"
		}
	}
	return result
}

func main() {
	// Check if the user enter one argument
	if len(os.Args[1:]) == 1 {
		if os.Args[1] == "" {
			return
		}
		// Read from the file standard
		data := tools.Read_Input("standard.txt")
		// Split by newline and after that delete the empty strings to organise the file
		elements := strings.Split(string(data[1:]), "\n")
		// Split the argument by new line to check every one
		result := DrawAsciiArt(tools.RemoveEmptyString(elements))
		// handling the additionnel new line if the arguiment is a bunche of new lines
		if tools.IsAllNl(result) {
			result = result[1:]
		}
		// Printing final result
		fmt.Printf(result)
	} else {
		log.Fatalln("Usage: go run . [STRING]\nEX: go run . \"something\"")
	}
}
