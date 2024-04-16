package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func replaceHex(input string) string {
	words := strings.Fields(input)
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			index, _ := strconv.ParseInt(words[i-1], 16, 64)
			words[i-1] = fmt.Sprintf("%d", index)
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}

func replaceBin(sample string) string {
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

func upper(input string) string {
	words := strings.Split(input, " ")
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "(up") {
			if strings.Contains(words[i], "(up,") {
				num, _ := strconv.Atoi(words[i+1][:len(words[i+1])-1])
				for j := i - num; j < i; j++ {
					words[j] = strings.ToUpper(words[j])
				}
				words = append(words[:i], words[i+2:]...)
			} else {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
			}
		}
	}
	return strings.Join(words, " ")
}

func lower(input string) string {
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

func cap(input string) string {
	words := strings.Split(input, " ")
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "(cap") {
			if strings.Contains(words[i], "(cap,") {
				num, _ := strconv.Atoi(words[i+1][:len(words[i+1])-1])
				for j := i - num; j < i; j++ {
					words[j] = strings.Title(words[j])
				}
				words = append(words[:i], words[i+2:]...)
			} else {
				words[i-1] = strings.Title(words[i-1])
				words = append(words[:i], words[i+1:]...)
			}
		}
	}
	return strings.Join(words, " ")
}

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
	return strings.TrimSpace(input)
}

func marks(text string) string {
	parts := strings.Split(text, "'")
	for i := 1; i < len(parts)-1; i += 2 {
		parts[i] = strings.TrimSpace(parts[i])
	}

	return strings.Join(parts, "'")
}

func vowel(input string) string {
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

func main() {
	inputFile := "sample.txt"
	outputFile := "result.txt"

	inputData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	input := string(inputData)

	formatted := replaceHex(input)
	modifiedVowels := replaceBin(formatted)
	quoted := upper(modifiedVowels)
	withNumbers := lower(quoted)
	cap := cap(withNumbers)
	mode := punctuation(cap)
	mark := marks(mode)
	finalOutput := vowel(mark)

	err = ioutil.WriteFile(outputFile, []byte(finalOutput), 0o644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(outputFile)
}
