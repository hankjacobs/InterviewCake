package main

import "fmt"

func main() {
	fmt.Println(generate("cats"))
}

func generate(str string) []string {
	fmt.Println("1: " + str)
	if len(str) <= 1 {
		return []string{str}
	}

	allButLast := str[:len(str)-1]
	last := str[len(str)-1:]

	fmt.Printf("2: %s %s\n", allButLast, last)
	permsOfAllButLast := generate(allButLast)

	perms := []string{}
	for _, perm := range permsOfAllButLast {
		fmt.Printf("3: %s %s\n", perm, last)
		for i := range perm {
			newPerm := perm[:i] + last + perm[i:]
			fmt.Println("4: " + newPerm)
			perms = append(perms, newPerm)
		}
	}

	return perms
}
