package main

import (
	"fmt"
	"io/ioutil"

	"go-reloaded/reloaded"
)

// Define the main function that reads sample.txt and write re result on the result.txt file
func main() {
	inputFile := "sample.txt"
	outputFile := "result.txt"

	inputData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	input := string(inputData)

	formatted := reloaded.ReplaceHex(input)
	modifiedVowels := reloaded.ReplaceBin(formatted)
	quoted := reloaded.Upper(modifiedVowels)
	withNumbers := reloaded.Lower(quoted)
	cap := reloaded.Cap(withNumbers)
	mode := reloaded.Punctuation(cap)
	mark := reloaded.Marks(mode)
	finalOutput := reloaded.Vowel(mark)

	finalOutput += "\n"
	err = ioutil.WriteFile(outputFile, []byte(finalOutput), 0o644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
