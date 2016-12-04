package main

import "fmt"

func main() {
	fmt.Println(sort([]int{2, 3, 4, 5, 3, 2, 4, 4, 5}, 5))
}

func sort(unsorted []int, maxScore int) (sorted []int) {
	sorted = make([]int, 0, len(unsorted))
	buckets := make([]int, maxScore+1)

	for _, score := range unsorted {
		buckets[score]++
	}

	for score, count := range buckets {
		for i := 0; i < count; i++ {
			sorted = append(sorted, score)
		}
	}

	return sorted
}
