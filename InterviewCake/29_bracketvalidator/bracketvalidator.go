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
	fmt.Println(validate("{ [ ] ( ) }"))
	fmt.Println(validate("{ [ ( ] ) }"))
	fmt.Println(validate("{ [ }"))
	fmt.Println(validate("]]]"))
}

func validate(str string) bool {
	s := stack{container: []rune{}}

	for _, r := range str {
		switch r {
		case '(', '{', '[':
			s.push(r)
			break
		case ')', '}', ']':

			if s.size() == 0 {
				return false
			}

			opener := s.pop()
			switch {
			case r == ')' && opener != '(':
				return false
			case r == '}' && opener != '{':
				return false
			case r == ']' && opener != '[':
				return false
			}
		}
	}

	return s.size() == 0
}
