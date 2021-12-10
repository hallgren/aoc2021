package main

import (
	"fmt"

	"github.com/hallgren/aoc2021"
)

type Stack struct {
	s []rune
}

func (s *Stack) Push(v rune) {
	s.s = append(s.s, v)
}

func (s *Stack) Pop() rune {
	v := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return v
}

func Point(r rune) int {
	switch r {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}
	panic(r)
}

func main() {
	result := 0
	lines := aoc2021.Lines("input")
L:
	for _, line := range lines {
		stack := Stack{}
		for _, l := range line {
			switch l {
			case '(', '[', '{', '<':
				stack.Push(l)
			case ')':
				v := stack.Pop()
				if v != '(' {
					result += Point(l)
					continue L
				}
			case ']':
				v := stack.Pop()
				if v != '[' {
					result += Point(l)
					continue L
				}
			case '}':
				v := stack.Pop()
				if v != '{' {
					result += Point(l)
					continue L
				}
			case '>':
				v := stack.Pop()
				if v != '<' {
					result += Point(l)
					continue L
				}
			}
		}
	}
	fmt.Println(result)
}
