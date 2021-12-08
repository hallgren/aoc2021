package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hallgren/aoc2021"
)

func Lines(s []string) [][]string {
	res := [][]string{}
	for _, v := range s {
		_, part2 := Line(v)
		res = append(res, part2)
	}
	return res
}

func Line(s string) ([]string, []string) {
	parts := strings.Split(s, " | ")
	first := strings.Fields(parts[0])
	second := strings.Fields(parts[1])
	return first, second
}

func Count(s []string) int {
	sum := 0
	for _, v := range s {
		switch len(v) {
		case 2, 3, 4, 7:
			sum++
		}
	}
	return sum
}

func CountAll(s [][]string) int {
	res := 0
	for _, v := range s {
		count := Count(v)
		res += count
	}
	return res
}

type Map struct {
	A string
	B int
	C int
	D int
	E int
	F int
	G int
}

type MapVal struct {
	Zero      string
	One       string // count 2
	Two       string
	Three     string // count 5 and include One
	Four      string // count 4
	Five      string
	Six       string
	Seven     string // count 3
	Eight     string // count 5
	Nine      string
	FiveCount []string
	SixCount  []string
}

func ThreeTwoFive(one, four string, options []string) (string, string, string) {
	var three string
	var two string
	var five string
	var res []string
	for _, option := range options {
		if strings.Contains(option, string(one[0])) && strings.Contains(option, string(one[1])) {
			three = option
		} else {
			res = append(res, option)
		}
	}

	for _, option := range res {
		sum := 0
		if strings.Contains(option, string(four[0])) {
			sum++
		}
		if strings.Contains(option, string(four[1])) {
			sum++
		}
		if strings.Contains(option, string(four[2])) {
			sum++
		}
		if strings.Contains(option, string(four[3])) {
			sum++
		}
		if sum == 2 {
			two = option
		} else {
			five = option
		}
	}
	return three, two, five
}

func SixNineZero(one, four string, options []string) (string, string, string) {
	var six string
	var nine string
	var zero string
	var res []string
	for _, option := range options {
		if strings.Contains(option, string(one[0])) && strings.Contains(option, string(one[1])) {
			res = append(res, option)
		} else {
			six = option
		}
	}
	for _, option := range res {
		sum := 0
		if strings.Contains(option, string(four[0])) {
			sum++
		}
		if strings.Contains(option, string(four[1])) {
			sum++
		}
		if strings.Contains(option, string(four[2])) {
			sum++
		}
		if strings.Contains(option, string(four[3])) {
			sum++
		}
		if sum == 4 {
			nine = option
		} else {
			zero = option
		}
	}
	return six, nine, zero
}

func Decode1(s []string) MapVal {
	mv := MapVal{}
	for _, val := range s {
		if len(val) == 2 {
			mv.One = SortStringByCharacter(val)
		} else if len(val) == 3 {
			mv.Seven = SortStringByCharacter(val)
		} else if len(val) == 4 {
			mv.Four = SortStringByCharacter(val)
		} else if len(val) == 7 {
			mv.Eight = SortStringByCharacter(val)
		} else if len(val) == 5 {
			mv.FiveCount = append(mv.FiveCount, SortStringByCharacter(val))
		} else if len(val) == 6 {
			mv.SixCount = append(mv.SixCount, SortStringByCharacter(val))
		}
	}
	mv.Three, mv.Two, mv.Five = ThreeTwoFive(mv.One, mv.Four, mv.FiveCount)
	mv.Six, mv.Nine, mv.Zero = SixNineZero(mv.One, mv.Four, mv.SixCount)
	return mv
}

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func Calc(mv MapVal, values []string) int {
	sum := 0
	for i, v := range values {
		p := 0
		switch SortStringByCharacter(v) {
		case mv.Zero:
			p = 0
		case mv.One:
			p = 1
		case mv.Two:
			p = 2
		case mv.Three:
			p = 3
		case mv.Four:
			p = 4
		case mv.Five:
			p = 5
		case mv.Six:
			p = 6
		case mv.Seven:
			p = 7
		case mv.Eight:
			p = 8
		case mv.Nine:
			p = 9
		}

		if i == 0 {
			sum += 1000 * p
		} else if i == 1 {
			sum += 100 * p
		} else if i == 2 {
			sum += 10 * p
		} else if i == 3 {
			sum += p
		}
	}
	return sum
}

func main() {
	lines := aoc2021.Lines("../input")
	sum := 0
	for _, line := range lines {
		f, s := Line(line)
		mv := Decode1(f)
		sum += Calc(mv, s)
	}
	fmt.Println(sum)

}
