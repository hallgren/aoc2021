package main

import (
	"fmt"
	"strings"

	"github.com/hallgren/aoc2021"
)

func Fishes(s string) [9]int {
	res := [9]int{}
	v := strings.Split(s, ",")
	for _, f := range v {
		res[aoc2021.Int(f)] += 1
	}
	return res
}

func Proc(fish [9]int) [9]int {
	newFish := [9]int{}
	fmt.Println("fish", fish)
	for i := 8; i >= 0; i-- {
		fmt.Println(i, fish[i])
		if i > 0 {
			newFish[i-1] = fish[i]
		} else {
			newFish[8] = fish[0]
			newFish[6] += fish[0]
		}
	}
	fmt.Println("newfish", newFish)
	return newFish
}

func Len(fish [9]int) int {
	sum := 0
	for _, v := range fish {
		sum += v
	}
	return sum
}

func main() {
	lines := aoc2021.Lines("../input")
	fish := Fishes(lines[0])
	for i := 1; i <= 256; i++ {
		fish = Proc(fish)
	}
	fmt.Println("sum", Len(fish))
}
