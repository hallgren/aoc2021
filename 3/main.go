package main

import (
	"fmt"

	"github.com/hallgren/aoc2021"
)

type Bucket struct {
	Zero int
	One  int
}

func main() {
	var counter = make([]Bucket, 12)
	values := aoc2021.Lines("input.txt")
	for _, value := range values {
		for i, v := range value {
			if v == 48 {
				counter[i].Zero++
			} else {
				counter[i].One++
			}
		}
	}
	fmt.Println(counter)
	fmt.Println(gamma(counter) * epsilon(counter))
}

func gamma(s []Bucket) int {
	result := 0
	times := 1
	for i := 1; i <= len(s); i++ {
		if s[len(s)-i].One > s[len(s)-i].Zero {
			result += times
		}
		times *= 2
	}
	fmt.Println(result)
	return result
}

func epsilon(s []Bucket) int {
	result := 0
	times := 1
	for i := 1; i <= len(s); i++ {
		if s[len(s)-i].One < s[len(s)-i].Zero {
			result += times
		}
		times *= 2
	}
	fmt.Println(result)
	return result
}
