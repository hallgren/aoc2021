package main

import (
	"fmt"
	"strconv"

	"github.com/hallgren/aoc2021"
)

func main() {
	lines := aoc2021.Lines("input")
	var currentValue int
	var sum int
	for _, line := range lines {
		if n, err := strconv.Atoi(line); err == nil {
			if currentValue != 0 {
				if n > currentValue {
					sum++
				}
			}
			currentValue = n
		} else {
			panic("not a integer")
		}
	}
	fmt.Println(sum)
}
