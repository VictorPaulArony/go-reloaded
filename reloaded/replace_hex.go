package reloaded

import (
	"fmt"
	"strconv"
	"strings"
)

// This function replaces a word preceding "(hex)" with the hexadecimal value of that word.
// Parse the preceding word as hexadecimal
// Check if the current word is "(hex)" and if it's preceded by another word
// If parsing fails, print an error message
// Replace the preceding word with its decimal value
// Remove the "(hex)" word from the slice
// Decrement the index to avoid skipping the next word
// Join the words back into a single string

func ReplaceHex(input string) string {
	words := strings.Fields(input)
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			index, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				fmt.Println("INVALID CONVERTION", err)
			}
			words[i-1] = fmt.Sprintf("%d", index)
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}
