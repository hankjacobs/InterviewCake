package main

import (
	"fmt"
)

func main() {
	var str = "find you will pain only go you recordings security the into if"
	bytes := []byte(str)
	reverseWords(bytes)

	rev := string(bytes[:])

	fmt.Println(rev)
}

func reverseWords(sentence []byte) {
	reverseCharacters(sentence)
	fmt.Println(string(sentence[:]))
	startOfWord := 0
	for i := 0; i <= len(sentence); i++ {
		if i == len(sentence) || sentence[i] == byte(' ') {
			reverseCharacters(sentence[startOfWord:i])
			startOfWord = i + 1
		}
	}
}

func reverseCharacters(chars []byte) {
	i := 0
	y := len(chars) - 1
	for i < y {
		chars[i], chars[y] = chars[y], chars[i]
		i++
		y--
	}
}
