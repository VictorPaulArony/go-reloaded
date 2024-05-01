package main

import (
	"fmt"
	"strconv"
	"strings"
)


func punctuation(input string) string {
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
	for _, group := range groups {
		input = strings.ReplaceAll(input, " "+group, group)
		input = strings.ReplaceAll(input, group+" ", group)
	}

	return strings.TrimSpace(input)
}

func vowels(input string) string {
	words := strings.Fields(input)
	vowels := "aeiouhAEIOUH"

	for i := 0; i < len(words)-1; i++ {
		switch {
		case words[i] == "(bin)" && i > 0:
			index, _ := strconv.ParseInt(words[i-1], 16, 64)
			words[i-1] = fmt.Sprintf("%d", index)
			words = append(words[:i], words[i+1:]...)
			i--
		case words[i] == "(hex)" && i > 0:
			index, _ := strconv.ParseInt(words[i-1], 16, 64)
			words[i-1] = fmt.Sprintf("%d", index)
			words = append(words[:i], words[i+1:]...)
			i--
		case words[i] == "(low)" && i > 0:
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		case words[i] == "(up)" && i > 0:
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		case strings.ContainsAny(words[i], vowels) && i > 0:
			if words[i] == "a" || words[i] == "A" {
				if strings.ContainsAny(string(words[i+1][0]), vowels) {
					words[i] = "an"
				}
			} else if words[i] == "an" || words[i] == "An" {
				if !strings.ContainsAny(string(words[i+1][0]), vowels) {
					words[i] = "a"
				}
			}
		}
	}

	return strings.Join(words, " ")
}

func marks(text string) string {
	parts := strings.Split(text, "'")
	for i := 1; i < len(parts)-1; i += 2 {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return strings.Join(parts, "'")
}

func main() {
	input := "This is a MANGO (low) example (up) an 1EF (hex) sentence. It ... contains a hour ,some (low) modifications. and 101 (bin)"
	formatted := punctuation(input)
	modified := vowels(formatted)
	quoted := marks(modified)
	fmt.Println(quoted)
}
