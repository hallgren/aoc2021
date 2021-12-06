package main

import (
	"fmt"
	"strings"

	"github.com/hallgren/aoc2021"
)

func Fishes(s string) []int {
	res := []int{}
	v := strings.Split(s, ",")
	for _, f := range v {
		res = append(res, aoc2021.Int(f)+1)
	}
	return res
}

func Proc(fish []int) []int {
	newFish := []int{}
	for i, f := range fish {
		f--
		if f < 7 && f%7 == 0 {
			newFish = append(newFish, 9)
			f = 7
		}
		fish[i] = f
	}
	fish = append(fish, newFish...)
	return fish
}

func Day(fish []int) []int {
	var divided [][]int

	/*
		chunkSize := (len(fish) + 8 - 1) / 8

			for i := 0; i < len(fish); i += chunkSize {
				end := i + chunkSize

				if end > len(fish) {
					end = len(fish)
				}

				divided = append(divided, fish[i:end])
			}*/
	res := []int{}
	for _, l := range divided {
		Print(l)
		res = append(res, Proc(l)...)
	}
	return Proc(fish)
}

func Print(fish []int) {
	r := []int{}
	for _, f := range fish {
		r = append(r, f-1)
	}
	fmt.Println(r)
}

func main() {
	lines := aoc2021.Lines("sample")
	fish := Fishes(lines[0])
	for i := 1; i <= 80; i++ {
		fish = Day(fish)
		fmt.Println(i, len(fish))
	}
}
