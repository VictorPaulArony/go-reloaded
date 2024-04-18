package main

import (
	"fmt"
	"strconv"
	"strings"
)
    // func replaceHex (input string) string {
    //     words := strings.Fields(input)
    //     for i:= 0; i < len(words); i++ {
    //         if words[i] == "(hex)" && i > 0 {
    //             index, _ := strconv.ParseInt(words[i-1], 16, 64)
    //             words[i-1] = fmt.Sprintf("%d", index)
    //             words = append(words[:i], words[i+1:]...)
    //             i--
    //         }
    //     }
    //     return strings.Join(words, " ")
    // }
	    func replaceHex(sample string) string {
		    words := strings.Fields(sample)
		    for i , word := range words {
			    if strings.Contains(word, "(hex)" ) {
				    num, _ := strconv.ParseInt(words[i-1], 16, 64)
				    words[i-1] = fmt.Sprintf("%d", num)
				    words = append(words[:i], words[i+1:]...)
				    }
			    }
		    return strings.Join(words, " ")
		    }
		    
    func main() {
        text := "look again in page 1E (hex)"
        output :=  replaceHex(text)
        fmt.Println(output)
        }
