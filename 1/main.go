package main

import (
	"fmt"

	"github.com/hallgren/aoc2021"
)

func main() {
	values := aoc2021.Ints("input")
	var currentValue int
	var sum int
	for _, value := range values {
		if currentValue != 0 {
			if value > currentValue {
				sum++
			}
		}
		currentValue = value
	}
	fmt.Println(sum)
}
