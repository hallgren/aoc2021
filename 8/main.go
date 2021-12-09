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

func main() {
	lines := aoc2021.Lines("input")
	f, s := Line(lines[0])
	l := Lines(lines)
	fmt.Println(f, s, l, CountAll(l))

}
