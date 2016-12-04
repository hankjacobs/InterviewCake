package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand7())
	fmt.Println(rand7())
	fmt.Println(rand7())
	fmt.Println(rand7())
	fmt.Println(rand7())
}

func rand7() int {
	num := (rand5()-1)*5 + (rand5() - 1) + 1 // [1, 25]
	for num > 21 {                           // constrain it to [1, 21] so all values in [1, 7] have equal probability of being generated
		num = (rand5()-1)*5 + (rand5() - 1) + 1
	}

	return num%7 + 1
}

func rand5() int {
	return rand.Intn(5) + 1
}
