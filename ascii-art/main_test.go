package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"

	"ascii-art/tools"
)

func AsciiArtBase(elements []string, something string) string {
	var result string
	words := strings.Split(something, `\n`)
	for _, word := range words {
		if word != "" {
			for j := 0; j < 8; j++ {
				for _, char := range word {
					if char < 32 || char > 126 {
						log.Fatalln("Error: please provide printable characters!!\nhelp: man ascii")
					} else {
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

func ascii_art(something string) string {
	var result string
	if something == "" {
		return result
	}
	data := tools.Read_Input("standard.txt")
	elements := strings.Split(string(data[1:]), "\n")
	result = AsciiArtBase(tools.RemoveEmptyString(elements), something)
	if tools.IsAllNl(result) {
		result = result[1:]
	}
	return result
}

func LoadTests() []string {
	testfile, err := os.Open("tests_input.txt")
	tools.CheckError(err, "Error opening testfile: \"tests_input.txt\"")
	defer testfile.Close()

	var tests []string
	scanner := bufio.NewScanner(testfile)
	for scanner.Scan() {
		tests = append(tests, scanner.Text())
	}
	tools.CheckError(scanner.Err(), "scanner error")
	return tests
}

func Test_main(t *testing.T) {
	//Read the file test and check it .
	tests := LoadTests()
	for _, test := range tests {
		//run the tests in the main and stock the result from stdout.
		got, err := exec.Command("go", "run", ".", test).Output()
		tools.CheckError(err,"")
		want := ascii_art(test)
		//Compare the result that the main.go give and the test give if they are the same
		if want == string(got) {
			t.Logf(test)
		} else {
			t.Fatal(test)
		}
	}
}
