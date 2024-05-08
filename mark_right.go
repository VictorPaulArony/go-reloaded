package main

import (
	"fmt"
	"strings"
)

func marks(input []string) []string {
	count := 0
	for i, char := range input {
		if char == "'" && count == 0 {
			count++
			input[i+1] = char + input[i+1]
			input = append(input[:i], input[i+1:]...)
		}
	}
	for i, char := range input {
		if char == "'\"" || char == "'" && count != 0 {
			count--
			input[i-1] = input[i-1] + char
			input = append(input[:i], input[i+1:]...)
		}
	}
	return input
}

func main() {
	sampleText := "I am exactly how they describe me: ' awesome '"
	words := strings.Fields(sampleText)
	formattedText := marks(words)
	fmt.Println(strings.Join(formattedText, " "))

	sampleText2 := "\"As Elton John said: ' I am the most well-known homosexual in the world   '\""
	words2 := strings.Fields(sampleText2)
	formattedText2 := marks(words2)
	fmt.Println(strings.Join(formattedText2, " "))
}
