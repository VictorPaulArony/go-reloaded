package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
// func replaceBin(sample string) string {
// 	words := strings.Fields(sample) 
// 	for i, word := range words {
// 		if strings.Contains(words[i], "(cap)") {
// 			bin, _ := strconv.ParseInt(wordsi-1], 2, 64 ) 
// 		}
// 		words[i-1] = fmt.Sprintf("%d", bin)
// 		words[i] = ""
// 	}
// }
return strings.Join(words, " " )
func main() {
	text := "It has been 01 (bin) years"
	transformed := replaceBin(text)
	fmt.Println(transformed)
}
