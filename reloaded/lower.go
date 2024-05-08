package reloaded

import (
	"strconv"
	"strings"
)

// Define a function to convert the input string to lowercase
func Lower(input string) string {
	words := strings.Split(input, " ")
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "(low") {
			if strings.Contains(words[i], "(low,") {
				num, _ := strconv.Atoi(words[i+1][:len(words[i+1])-1])
				for j := i - num; j < i; j++ {
					words[j] = strings.ToLower(words[j])
				}
				words = append(words[:i], words[i+2:]...)
			} else {
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...)
			}
		}
	}
	return strings.Join(words, " ")
}
