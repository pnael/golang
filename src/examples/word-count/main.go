package main

import (
	"fmt"
	"strings"
)

func WordCount(words string) map[string]int {
	//Count the number of word in words, retrun a map
	tabwords := strings.Fields(words)
	result := map[string]int{}

	for _, word := range tabwords {
		result[word]++
	}
	return result
}

func main() {
	words := "Bonjour, Noël est passé. Encore un mot est et un"
	result := WordCount(words)

	fmt.Println(result)
}
