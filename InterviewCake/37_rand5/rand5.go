package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand5())
	fmt.Println(rand5())
	fmt.Println(rand5())
	fmt.Println(rand5())
	fmt.Println(rand5())
}

func rand5() int {

	num := rand7()
	for num > 5 {
		num = rand7()
	}

	return num
}

func rand7() int {
	return rand.Intn(7) + 1
}
