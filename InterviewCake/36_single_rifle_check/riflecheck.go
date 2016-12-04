package main

import "fmt"

func main() {
	half1 := []int{1, 2, 3, 4, 5}
	half2 := []int{6, 7, 8, 9, 10}
	riffleShuffled := []int{1, 6, 7, 2, 3, 8, 4, 5, 9, 10}
	randShuffled := []int{1, 6, 9, 2, 3, 8, 4, 5, 7, 10}

	fmt.Printf("riffleShuffled: %t\n", rifleCheck(riffleShuffled, half1, half2))
	fmt.Printf("randShuffled: %t\n", rifleCheck(randShuffled, half1, half2))
}

func rifleCheck(shuffled []int, half1 []int, half2 []int) bool {

	t1 := 0
	t2 := 0
	for _, val := range shuffled {
		if t1 < len(half1) && half1[t1] == val {
			t1++
		} else if t2 < len(half2) && half2[t2] == val {
			t2++
		} else {
			return false
		}
	}

	return true
}
