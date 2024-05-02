package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "There it was. A amazing rock and an pen! An pilot is good"
	fmt.Println("Original:", input)
	transformed := vowel(input)
	fmt.Println("Transformed:", transformed)
}
func vowel(input string) string {
    words := strings.Fields(input) 
    vowels := "aeiouhAEIOUH"
    for i := 0; i < len(words); i++ {
        if words[i] == "a" || words[i] == "A" && strings.ContainsAny(string(words[i+1][0]), vowels) {
            if words[i] == "a"{
            words[i] = "an"
        } else {
            words[i] = "An"
        }
    }
    if words[i] == "an" || words[i] == "An" && !strings.ContainsAny(string(words[i+1][0]), vowels) {
        if words[i] == "an" {
            words[i] = "a"
        } else {
            words[i] = "A"
        }
    }
    }
    return strings.Join(words, " ")
}
