/*Write a function that takes a string of text (a sentence) and returns a map where the keys are the words and the values are the count of how many times each word appears*/

package main
import (
	"fmt"
	"strings"
)

func wordCount(text string) map[string]int {
	wordMap:=make(map[string]int)
	words := strings.Fields(text)

	for _, w := range words {
		wordMap[w]++
	}
	return wordMap
}

func main() {
	sentence := "Go is an open source programming language that makes it easy to build simple reliable and efficient software Go is expressive concise clean and efficient"
	result := wordCount(sentence)
	fmt.Println("Word Count Map:", result)
}