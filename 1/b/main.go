package main

import (
	"fmt"

	"github.com/hallgren/aoc2021"
)

func main() {
	values := aoc2021.Ints("../input")
	var sum int
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
