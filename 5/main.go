package main

import (
	"fmt"
	"strings"

	"github.com/hallgren/aoc2021"
)

type Point struct {
	X int
	Y int
}

type Line struct {
	Start Point
	Stop  Point
}

func (l Line) Points() []Point {
	points := []Point{}
	x := Range(Min(l.Start.X, l.Stop.X), Max(l.Start.X, l.Stop.X))
	y := Range(Min(l.Start.Y, l.Stop.Y), Max(l.Start.Y, l.Stop.Y))
	for _, x1 := range x {
		for _, y1 := range y {
			points = append(points, Point{X: x1, Y: y1})
		}
	}
	return points
}

func Range(start, stop int) []int {
	r := []int{}
	for ; start <= stop; start++ {
		r = append(r, start)
	}
	return r
}

func NewLine(s string) *Line {
	parts := strings.Fields(s)
	first := strings.Split(parts[0], ",")
	second := strings.Split(parts[2], ",")
	start := Point{X: aoc2021.Int(first[0]), Y: aoc2021.Int(first[1])}
	stop := Point{X: aoc2021.Int(second[0]), Y: aoc2021.Int(second[1])}

	if start.X != stop.X && start.Y != stop.Y {
		return nil
	}

	line := Line{
		Start: start,
		Stop:  stop,
	}
	return &line
}

func NewLines(l []string) []Line {
	r := []Line{}
	for _, l1 := range l {
		line := NewLine(l1)
		if line != nil {
			r = append(r, *line)
		}
	}
	return r
}

func main() {
	m := make(map[string]int)
	l := aoc2021.Lines("input")
	lines := NewLines(l)
	for _, line := range lines {
		for _, point := range line.Points() {
			m[fmt.Sprintf("%d-%d", point.X, point.Y)] += 1
		}
	}
	sum := 0
	for k, v := range m {
		if v >= 2 {
			fmt.Println(k)
			sum++
		}
	}
	fmt.Println(sum)
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
