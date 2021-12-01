package main

import (
	"fmt"
	"strconv"

	"github.com/hallgren/aoc2021"
)

func main() {
	lines := aoc2021.Lines("../input")
	var values []int
	var sum int
	for _, line := range lines {
		if n, err := strconv.Atoi(line); err == nil {
			values = append(values, n)
		} else {
			panic("not a integer")
		}
	}
	for i, _ := range values {
		if i >= 3 {
			a := values[i-1] + values[i-2] + values[i-3]
			b := values[i] + values[i-1] + values[i-2]
			if b > a {
				sum++
			}
		}
	}
	fmt.Println(sum)
}
