package main

import (
	"fmt"
	"strings"

	"github.com/hallgren/aoc2021"
)

func Lines(s []string) [][]string {
	res := [][]string{}
	for _, v := range s {
		_, part2 := Line(v)
		res = append(res, part2)
	}
	return res
}

func Line(s string) ([]string, []string) {
	parts := strings.Split(s, " | ")
	first := strings.Fields(parts[0])
	second := strings.Fields(parts[1])
	return first, second
}

func Count(s []string) int {
	sum := 0
	for _, v := range s {
		switch len(v) {
		case 2, 3, 4, 7:
			sum++
		}
	}
	return sum
}

func CountAll(s [][]string) int {
	res := 0
	for _, v := range s {
		count := Count(v)
		res += count
	}
	return res
}

type Map struct {
	A string
	B int
	C int
	D int
	E int
	F int
	G int
}

type MapVal struct {
	One       string
	Four      string
	Seven     string
	Eight     string
	FiveCount []string
	SixCount  []string
}

func Decode1(s []string) MapVal {
	mv := MapVal{}
	for _, val := range s {
		if len(val) == 2 {
			fmt.Println(val)
			mv.One = val
		} else if len(val) == 3 {
			fmt.Println(val)
			mv.Seven = val
		} else if len(val) == 4 {
			mv.Four = val
		} else if len(val) == 7 {
			mv.Eight = val
		} else if len(val) == 5 {
			mv.FiveCount = append(mv.FiveCount, val)
		} else if len(val) == 6 {
			mv.SixCount = append(mv.SixCount, val)
		}
	}
	return mv
}

func A(one, seven string) string {
	for _, v := range seven {
		for _, b := range one {
			if b != v {
				return string(v)
			}
		}
	}
	panic("error")
}

func C(one string, fives []string) string {
	return ""
}

func Decode2(mv MapVal) Map {
	m := Map{}
	m.A = A(mv.One, mv.Seven)
	fmt.Println(m)
	return m
}

func main() {
	lines := aoc2021.Lines("sample")
	f, s := Line(lines[0])
	mv := Decode1(f)
	m := Decode2(mv)

	fmt.Println(f, s, mv, m)

}
