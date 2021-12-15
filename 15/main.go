package main

import (
	"fmt"

	"github.com/hallgren/aoc2021"
)

const Size = 10

type Matrix [Size][Size]int

type Point struct {
	X     int
	Y     int
	Value int
}

type Path struct {
	Points []Point
}

func (p Point) Valid(m Matrix) bool {
	if p.Y < 0 || p.Y >= len(m[0]) {
		return false
	}
	if p.X < 0 || p.X >= len(m) {
		return false
	}
	return true
}

func (p Point) Neighbours(m Matrix) []Point {
	res := []Point{}
	// above
	po := Point{X: p.X, Y: p.Y - 1}
	if po.Valid(m) {
		po.Value = m[po.Y][po.X]
		res = append(res, po)
	}
	// right
	po = Point{X: p.X + 1, Y: p.Y}
	if po.Valid(m) {
		po.Value = m[po.Y][po.X]
		res = append(res, po)
	}
	// under
	po = Point{X: p.X, Y: p.Y + 1}
	if po.Valid(m) {
		po.Value = m[po.Y][po.X]
		res = append(res, po)
	}
	// left
	po = Point{X: p.X - 1, Y: p.Y}
	if po.Valid(m) {
		po.Value = m[po.Y][po.X]
		res = append(res, po)
	}
	return res
}

func CreateM(lines []string) Matrix {
	m := Matrix{}
	for y, line := range lines {
		for x, value := range line {
			m[y][x] = aoc2021.Int(string(value))
		}
	}
	return m
}

func (p Point) End(m Matrix) bool {
	fmt.Println("end", p, len(m), len(m[0]))
	if p.X == len(m)-1 && p.Y == len(m[0])-1 {
		return true
	}
	return false
}

func Run(path []Point, point Point, m Matrix, res []Path) {
	path = append(path, point)
	if point.End(m) {
		fmt.Println("last point", point)
		p := Path{path}
		res = append(res, p)
		return
	} else {
		neighbours := point.Neighbours(m)
		fmt.Println("neighbours", neighbours)
		for _, n := range neighbours {
			if n.Valid(m) {
				Run(path, n, m, res)
			}
		}

	}
}

func main() {
	lines := aoc2021.Lines("sample")
	m := CreateM(lines)
	start := Point{X: 0, Y: 0, Value: 0}
	res := []Path{}
	Run([]Point{}, start, m, res)
	fmt.Println(res, start.Neighbours(m))

}
