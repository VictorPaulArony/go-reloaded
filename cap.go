package main

import (
	"fmt"
	"strings"
)
    func cap(sample string) string {
        words := strings.Fields(sample)
        for i := 0; i < len(words)-1; i++ {
            if words[i] == "(cap)" && i > 0 {
                words[i-1] = strings.Title(words[i-1])
                words = append(words[:i], words[i+1:]...)
                i++
            }
        }
        return strings.Join(words, " ")
    }
    func main() {
        text := "The quick brown fox jumped (cap) over the lazy dog"
        result := cap(text)
        fmt.Println(result)
    }