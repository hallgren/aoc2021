package main

import (
	"fmt"
	"sort"

	"github.com/hallgren/aoc2021"
)

type Stack struct {
	s []rune
}

func (s *Stack) Push(v rune) {
	s.s = append(s.s, v)
}

func (s *Stack) Pop() (rune, bool) {
	if len(s.s) == 0 {
		return 'a', false
	}
	v := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return v, true
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

func Incomplete() []Stack {
	incomplete := []Stack{}
	lines := aoc2021.Lines("../input")
L:
	for _, line := range lines {
		stack := Stack{}
		for _, l := range line {
			switch l {
			case '(', '[', '{', '<':
				stack.Push(l)
			case ')':
				v, _ := stack.Pop()
				if v != '(' {
					continue L
				}
			case ']':
				v, _ := stack.Pop()
				if v != '[' {
					continue L
				}
			case '}':
				v, _ := stack.Pop()
				if v != '{' {
					continue L
				}
			case '>':
				v, _ := stack.Pop()
				if v != '<' {
					continue L
				}
			}
		}
		incomplete = append(incomplete, stack)
	}
	return incomplete
}

func main() {
	allStackSums := []int{}
	incomplete := Incomplete()
	fmt.Println(incomplete)
	for _, s := range incomplete {
		stackSum := 0
		for {
			v, ok := s.Pop()
			if !ok {
				break
			}
			stackSum *= 5
			switch v {
			case '(':
				stackSum += 1
			case '[':
				stackSum += 2
			case '{':
				stackSum += 3
			case '<':
				stackSum += 4
			}
		}
		allStackSums = append(allStackSums, stackSum)
	}
	sort.Ints(allStackSums)
	v := len(allStackSums) / 2
	fmt.Println(allStackSums, v, allStackSums[v], allStackSums[v+1])
}
