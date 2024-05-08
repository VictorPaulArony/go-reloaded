package reloaded

import "strings"

// Define a function to convert the input string to vowel form
func Punctuation(input string) string {
	punctuations := []string{".", ",", "!", "?", ":", ";"}
	groups := []string{"....", "!?"}

	for _, group := range groups {
		input = strings.ReplaceAll(input, " "+group+" ", group)
		input = strings.ReplaceAll(input, group+" ", group)
		input = strings.ReplaceAll(input, " "+group, group)
	}
	for _, punct := range punctuations {
		input = strings.ReplaceAll(input, punct+" ", punct)
		input = strings.ReplaceAll(input, punct, punct+" ")
		input = strings.ReplaceAll(input, " "+punct, punct)
	}
	return strings.TrimSpace(input)
}
