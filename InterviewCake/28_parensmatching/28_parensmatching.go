package main

import "fmt"

type stack struct {
	container []rune
}

func (s *stack) push(r rune) {
	s.container = append(s.container, r)
}

func (s *stack) pop() rune {
	top := len(s.container) - 1
	r := s.container[top]
	s.container = s.container[:top]
	return r
}

func (s *stack) size() int {
	return len(s.container)
}

func main() {
	str := "Sometimes (when I nest them (my parentheticals) too much (like this (and this))) they get confusing."
	matched1 := match1(10, str)
	matched2 := match2(10, str)

	fmt.Printf("%d %d", matched1, matched2)
}

func match1(startpos int, str string) int {
	i := startpos
	s := stack{container: []rune{}}

	for ; i < len(str); i++ {
		r := str[i]
		if r == '(' {
			s.push(rune(r))
		} else if r == ')' {
			s.pop()
		}

		if s.size() == 0 {
			break
		}
	}

	return i
}

func match2(startpos int, str string) int {
	i := startpos
	count := 0
	for ; i < len(str); i++ {
		r := str[i]
		if r == '(' {
			count++
		} else if r == ')' {
			count--
		}

		if count == 0 {
			break
		}
	}

	return i
}
