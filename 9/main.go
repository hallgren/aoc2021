package main

import (
	"fmt"

	"github.com/hallgren/aoc2021"
)

func Lowest(p Point, points []Point) bool {
	for _, v := range points {
		if v.Value <= p.Value {
			return false
		}
	}
	return true
}

type Point struct {
	X     int
	Y     int
	Value int
}

type Row struct {
	Points []Point
}

type Matrix struct {
	Rows []Row
}

func (m Matrix) MaxX() int {
	return len(m.Rows[0].Points)
}

func (m Matrix) MaxY() int {
	return len(m.Rows)
}

func (m Matrix) Point(x, y int) Point {
	return m.Rows[y].Points[x]
}

func ValidPoint(p Point, m Matrix) bool {
	if p.X < 0 || p.X >= m.MaxX() {
		return false
	}
	if p.Y < 0 || p.Y >= m.MaxY() {
		return false
	}
	return true
}

func ValidPoints(p Point, matrix Matrix) []Point {
	res := []Point{}
	//above
	above := Point{X: p.X, Y: p.Y - 1}
	if ValidPoint(above, matrix) {
		res = append(res, matrix.Point(above.X, above.Y))
	}
	// right
	right := Point{X: p.X + 1, Y: p.Y}
	if ValidPoint(right, matrix) {
		res = append(res, matrix.Point(right.X, right.Y))
	}
	// left
	left := Point{X: p.X - 1, Y: p.Y}
	if ValidPoint(left, matrix) {
		res = append(res, matrix.Point(left.X, left.Y))
	}
	// down
	down := Point{X: p.X, Y: p.Y + 1}
	if ValidPoint(down, matrix) {
		res = append(res, matrix.Point(down.X, down.Y))
	}
	return res
}

func CreateMatrix(lines []string) Matrix {
	matrix := Matrix{}
	for i, line := range lines {
		row := []Point{}
		for j, e := range line {
			v := aoc2021.Int(string(e))
			p := Point{Y: i, X: j, Value: v}
			row = append(row, p)
		}
		matrix.Rows = append(matrix.Rows, Row{Points: row})
	}
	return matrix
}

func (m Matrix) Sum() int {
	sum := 0
	for _, r := range m.Rows {
		for _, p := range r.Points {
			po := ValidPoints(p, m)
			if Lowest(p, po) {
				sum += p.Value + 1
			}
		}
	}
	return sum
}

func main() {
	lines := aoc2021.Lines("input")
	m := CreateMatrix(lines)
	fmt.Println(m, m.Sum())
}
