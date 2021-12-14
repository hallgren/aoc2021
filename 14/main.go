package main

import (
	"fmt"
	"strings"

	"github.com/hallgren/aoc2021"
)

func CreateMap(lines []string) map[string]string {
	m := make(map[string]string)
	for _, line := range lines {
		p := strings.Split(line, " ")
		m[p[0]] = p[2]
	}
	return m
}

func Run(template string, m map[string]string) string {
	res := ""
	first, last := "", ""
	for _, t := range template {
		first = last
		last = string(t)
		if last != "" {
			v := m[first+last]
			res += first + v
		}
	}
	res += last
	return res
}

func HighLow(template string) (int, int) {
	high, low := 0, 1000000
	m := make(map[string]int)
	for _, t := range template {
		m[string(t)]++
	}
	for _, v := range m {
		if v > high {
			high = v
		}
		if v < low {
			low = v
		}
	}
	return high, low
}

func main() {
	lines := aoc2021.Lines("sample")
	template := lines[0]
	m := CreateMap(lines[2:])
	fmt.Println(template, m)
	for i := 0; i < 10; i++ {
		template = Run(template, m)
		fmt.Println(i, len(template))
	}
	h, l := HighLow(template)
	fmt.Println(h, l, h-l)
}
