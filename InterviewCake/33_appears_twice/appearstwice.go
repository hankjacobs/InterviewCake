package main

import "fmt"

func main() {
	fmt.Println(which([]int{1, 2, 3, 4, 4, 5, 6, 7, 8}, 8))
}

func which(numbers []int, n int) int {
	expected := (n*n + n) / 2

	got := 0
	for _, v := range numbers {
		got += v
	}

	return got - expected
}
