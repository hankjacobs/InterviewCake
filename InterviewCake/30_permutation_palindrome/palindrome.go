package main

import "fmt"

func main() {
	fmt.Println(check("civic"))
	fmt.Println(check("ivicc"))
	fmt.Println(check("civil"))
	fmt.Println(check("livci"))
	fmt.Println(check("lelel"))
	fmt.Println(check("lllee"))
}

func check(str string) bool {

	m := make(map[rune]int)

	for _, r := range str {
		m[r]++
	}

	evenCount := 0
	oddCount := 0
	for _, v := range m {
		if v%2 != 0 {
			oddCount++
		} else {
			evenCount++
		}
	}

	return oddCount <= 1
}
