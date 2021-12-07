package main

import (
	"fmt"
	"strings"

	"github.com/hallgren/aoc2021"
)

func Crabs(s string) []int {
	res := []int{}
	i := strings.Split(s, ",")
	for _, c := range i {
		res = append(res, aoc2021.Int(c))
	}
	return res
}

func Max(values []int) int {
	max := 0
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

func Min(values []int) int {
	min := 0
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

func HighLow(values []int) (int, int) {
	return Max(values), Min(values)
}

func Fuel(crabs []int, high, low int) int {
	res := 10000000
	for i := low; i <= high; i++ {
		iRes := 0
		for _, crab := range crabs {
			if crab > i {
				iRes += crab - i
			} else if crab < i {
				iRes += i - crab
			}
		}
		if res > iRes {
			res = iRes
		}
	}
	return res
}

func main() {
	lines := aoc2021.Lines("input")
	crabs := Crabs(lines[0])
	high, low := HighLow(crabs)
	res := Fuel(crabs, high, low)
	fmt.Println(crabs, high, low, res)
}
