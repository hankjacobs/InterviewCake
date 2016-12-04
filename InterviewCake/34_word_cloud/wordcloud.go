package main

import "strings"
import "fmt"

func main() {
	words := "We came, we saw, we conquered...then we ate Bill's (Mille-Feuille) cake."
	fmt.Println(wordCloud(words))
}

func wordCloud(words string) map[string]int {
	cloud := make(map[string]int)

	wordList := strings.FieldsFunc(words, func(r rune) bool {
		switch r {
		case ' ', '!', '.', '?', ';', ':', ',', '"', '\'', '(', ')':
			return true
		default:
			return false
		}
	})

	for _, word := range wordList {
		cloud[strings.ToLower(word)]++
	}

	return cloud
}
