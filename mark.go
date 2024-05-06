package main
import (
	"fmt"
	"strings"
)
// func marks(text string) string {
// 	parts := strings.Split(text, "'")
// 	for i := 1; i < len(parts)-1; i += 2 {
// 		parts[i] = strings.TrimSpace(parts[i])
// 	}
	
// 	return strings.Join(parts, "'")
// }
func marks(text string) string {
	marks := []string {"'"}
	for _, mark := range marks {
		text = strings.ReplaceAll(text, " "+mark, mark)
		text = strings.ReplaceAll(text, mark+" ", mark)
	}
	return strings.TrimSpace(text)
}
func main() {
	sampleText := "I am exactly how they describe me: ' awesome '"
	formattedText := marks(sampleText)
	fmt.Println(formattedText)
	sampleText2 := "\"As Elton John said: ' I am the most well-known homosexual in the world '\""
	formattedText2 := marks(sampleText2)
	fmt.Println(formattedText2)
}
