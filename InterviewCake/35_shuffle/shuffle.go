package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(shuffle([]int{1, 2, 3, 4}))
}

func random(start int, end int) int {
	if start == end {
		return start
	}

	rand.Seed(time.Now().UnixNano())
	return rand.Intn((end-start)+1) + start
}

func shuffle(deck []int) []int {

	for i := range deck {
		randomIndex := random(i, len(deck)-1)
		randomValue := deck[randomIndex]
		swapValue := deck[i]

		deck[i] = randomValue
		deck[randomIndex] = swapValue
	}

	return deck
}
