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

func Basin(p Point, points []Point) []Point {
	res := []Point{}
	for _, v := range points {
		if p.Value+1 == v.Value {
			res = append(res, v)
		}
	}
	return res
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

func (m Matrix) BasinStart(p Point) []Point {
	points := []Point{}
	po := ValidPoints(p, m)
	if Lowest(p, po) {
		points = append(points, p)
		for _, v := range po {
			if v.Value == p.Value+1 {
				points = append(points, v)
			}
		}
		fmt.Println("Lowest", p, points)
	}
	return points
}

func (m Matrix) Start(p Point) bool {
	po := ValidPoints(p, m)
	return Lowest(p, po)
}

func (m Matrix) B(p Point, vp []Point) []Point {
	res := []Point{}
	if p.Value == 8 {
		return res
	}
	for _, vpv := range vp {
		if p.Value+1 == vpv.Value {
			res = append(res, vpv)
			po := ValidPoints(vpv, m)
			r := m.B(vpv, po)
		L:
			for _, newPoint := range r {
				for _, resTemp := range res {
					if newPoint.X == resTemp.X && newPoint.Y == resTemp.Y {
						continue L
					}
				}
				res = append(res, newPoint)
			}
		}
	}
	return res
}

func (m Matrix) Go() {
	sums := []int{}
	for _, r := range m.Rows {
		for _, p := range r.Points {
			if m.Start(p) {
				points := []Point{p}
				po := ValidPoints(p, m)
				points = append(points, m.B(p, po)...)
				fmt.Println("Go points", p, points, len(points))
				sums = append(sums, len(points))
			}
		}
	}
	fmt.Println("Sums", sums)
	total := 1
	for i := 0; i < 3; i++ {
		highest := 0
		index := 0
		for j, s := range sums {
			if s > highest {
				highest = s
				index = j
			}
		}
		fmt.Println("highest", highest)
		total *= highest
		sums[index] = sums[len(sums)-1] // Copy last element to index i.
		sums[len(sums)-1] = 0           // Erase last element (write zero value).
		sums = sums[:len(sums)-1]
	}
	fmt.Println("Total", total)
}

func (m Matrix) SumBasin() int {
	points := []Point{}
	for _, r := range m.Rows {
		for _, p := range r.Points {
			points = append(points, m.BasinStart(p)...)
		}
	}
	fmt.Println("Total points ", len(points))
	return len(points)
}

func main() {
	lines := aoc2021.Lines("../input")
	m := CreateMatrix(lines)

	m.Go()
}
