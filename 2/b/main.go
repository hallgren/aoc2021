package main

import (
	"fmt"
	"strings"

	"github.com/hallgren/aoc2021"
)

func main() {
	values := aoc2021.Lines("../input")
	var hor int
	var dept int
	var aim int
	for _, value := range values {
		s := strings.Split(value, " ")
		value := aoc2021.Int(s[1])
		switch s[0] {
		case "forward":
			hor += value
			dept += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}
	fmt.Println(hor * dept)
}
