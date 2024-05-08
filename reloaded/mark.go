package reloaded

import "strings"

// Define a function to add marks to the input string
func Marks(input string) string {
	words := strings.Fields(input)
	count := 0
	for i, word := range words {
		if word == "'" && count == 0 {
			count++
			words[i+1] = "'" + words[i+1]
			words = append(words[:i], words[i+1:]...)
		}
	}
	for i, word := range words {
		if word == "'\"" || word == "'" && count != 0 {
			count--
			words[i-1] = words[i-1] + word
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")
}
