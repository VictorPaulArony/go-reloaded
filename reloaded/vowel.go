package reloaded

import (
	"strings"
)

// Define a function to convert the input string to vowel form
func Vowel(input string) string {
	words := strings.Fields(input)
	vowels := "aeiouhAEIOUH"
	for i := 0; i < len(words); i++ {
		if (words[i] == "a" || words[i] == "A") && strings.ContainsAny(string(words[i+1][0]), vowels) {
			if words[i] == "a" {
				words[i] = "an"
			} else {
				words[i] = "An"
			}
		}
		if (words[i] == "an" || words[i] == "An") && !strings.ContainsAny(string(words[i+1][0]), vowels) {
			if words[i] == "an" {
				words[i] = "a"
			} else {
				words[i] = "A"
			}
		}
	}
	return strings.Join(words, " ")
}
