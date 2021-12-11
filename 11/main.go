package main

import (
	"fmt"

	"github.com/hallgren/aoc2021"
)

type Octopus struct {
	X          int
	Y          int
	Value      int
	Flashed    bool
	Neighbours []*Octopus
}

func (o *Octopus) Add() {
	if !o.Flashed {
		o.Value++
		if o.Value > 9 {
			o.Value = 0
			o.Flashed = true
			for _, n := range o.Neighbours {
				n.Add()
			}
		}
	}
}

type Row struct {
	Octopus []*Octopus
}

type Matrix struct {
	Rows    []Row
	Flashes int
}

func (m Matrix) MaxX() int {
	return len(m.Rows[0].Octopus)
}

func (m Matrix) MaxY() int {
	return len(m.Rows)
}

func (m Matrix) Octopus(x, y int) *Octopus {
	for _, row := range m.Rows {
		for _, o := range row.Octopus {
			if x == o.X && y == o.Y {
				return o
			}
		}
	}
	return nil
}

func CreateMatrix(lines []string) Matrix {
	matrix := Matrix{}
	// Create octopus
	for i, line := range lines {
		row := []*Octopus{}
		for j, e := range line {
			v := aoc2021.Int(string(e))
			p := Octopus{Y: i, X: j, Value: v}
			row = append(row, &p)
		}
		matrix.Rows = append(matrix.Rows, Row{Octopus: row})
	}

	// Set neighbours
	for _, row := range matrix.Rows {
		for _, o := range row.Octopus {
			n := Neighbours(o, matrix)
			o.Neighbours = n
		}
	}
	return matrix
}

func Valid(p Octopus, m Matrix) bool {
	if p.X < 0 || p.X >= m.MaxX() {
		return false
	}
	if p.Y < 0 || p.Y >= m.MaxY() {
		return false
	}
	return true
}

func Neighbours(p *Octopus, matrix Matrix) []*Octopus {
	res := []*Octopus{}
	//aboveleft
	o := matrix.Octopus(p.X-1, p.Y-1)
	if o != nil {
		res = append(res, o)
	}
	//above
	o = matrix.Octopus(p.X, p.Y-1)
	if o != nil {
		res = append(res, o)
	}
	//above right
	o = matrix.Octopus(p.X+1, p.Y-1)
	if o != nil {
		res = append(res, o)
	}
	// left
	o = matrix.Octopus(p.X-1, p.Y)
	if o != nil {
		res = append(res, o)
	}
	// right
	o = matrix.Octopus(p.X+1, p.Y)
	if o != nil {
		res = append(res, o)
	}
	// down left
	o = matrix.Octopus(p.X-1, p.Y+1)
	if o != nil {
		res = append(res, o)
	}
	// down
	o = matrix.Octopus(p.X, p.Y+1)
	if o != nil {
		res = append(res, o)
	}
	// down right
	o = matrix.Octopus(p.X+1, p.Y+1)
	if o != nil {
		res = append(res, o)
	}
	return res
}

func (m *Matrix) Add() {
	for _, row := range m.Rows {
		for _, o := range row.Octopus {
			o.Add()
		}
	}

	// reset flash
	for _, row := range m.Rows {
		for _, o := range row.Octopus {
			if o.Flashed {
				o.Flashed = false
				m.Flashes++
			}
		}
	}

}

func main() {
	lines := aoc2021.Lines("input")
	matrix := CreateMatrix(lines)
	for _, row := range matrix.Rows {
		for _, o := range row.Octopus {
			fmt.Println(o.Value)
		}
		fmt.Println("-----")
	}
	for i := 0; i < 100; i++ {

		matrix.Add()
	}
	fmt.Println("-----")
	for _, row := range matrix.Rows {
		for _, o := range row.Octopus {
			fmt.Println(o.Value)
		}
		fmt.Println("-----")
	}
	fmt.Println("flashes ", matrix.Flashes)
}
