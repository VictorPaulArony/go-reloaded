package reloaded

import (
	"fmt"
	"strconv"
	"strings"
)

// Define a function to replace binary numbers in the input string
func ReplaceBin(sample string) string {
	words := strings.Split(sample, " ")
	for i, word := range words {
		if word == "(bin)" {
			binary, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				panic(err)
			}
			words[i-1] = fmt.Sprintf("%d", binary)
			words[i] = ""
		}
	}
	return strings.Join(words, " ")
}
