package main

import (
	"fmt"
	"strings"

	"github.com/hallgren/aoc2021"
)

const Size = 1500

var ResetRow = [Size]bool{}

type Matrix struct {
	M [Size][Size]bool
}

func CreateMatrix(lines []string) *Matrix {
	matrix := Matrix{}
	for _, line := range lines {
		if line == "" {
			break
		}
		p := strings.Split(line, ",")
		x := aoc2021.Int(p[0])
		y := aoc2021.Int(p[1])
		matrix.M[y][x] = true
	}
	return &matrix
}

func Merge(a, b [Size]bool) [Size]bool {
	for i := 0; i < Size; i++ {
		if a[i] || b[i] {
			a[i] = true
		}
	}
	return a
}

func (m *Matrix) FoldX(x int) {
	for i := 0; i < x; i++ {
		for y := 0; y < Size; y++ {
			left := m.M[y][x-1-i]
			right := m.M[y][x+1+i]
			if left || right {
				m.M[y][x-1-i] = true
			}
			m.M[y][x+1+i] = false
		}
	}
}

func (m *Matrix) FoldY(y int) {
	for i := 0; i < y; i++ {
		rowAbove := m.M[y-1-i]
		rowBelow := m.M[y+1+i]
		//fmt.Println(rowAbove)
		//fmt.Println(rowBelow)
		rowAbove = Merge(rowAbove, rowBelow)
		//fmt.Println(rowAbove)
		m.M[y-1-i] = rowAbove
		m.M[y+1+i] = ResetRow
		//fmt.Println(i, y-1-i, y+1+i, ".............")
	}
}

func (m *Matrix) Print() {
	for y := 0; y < 6; y++ {
		fmt.Println(m.M[y][:40])
	}
}

func (m *Matrix) Count() int {
	count := 0
	for _, row := range m.M {
		for _, p := range row {
			if p {
				count++
			}
		}

	}
	return count
}

func main() {
	lines := aoc2021.Lines("../input")
	matrix := CreateMatrix(lines)
	matrix.FoldX(655)
	matrix.FoldY(447)
	matrix.FoldX(327)
	matrix.FoldY(223)
	matrix.FoldX(163)
	matrix.FoldY(111)
	matrix.FoldX(81)
	matrix.FoldY(55)
	matrix.FoldX(40)
	matrix.FoldY(27)
	matrix.FoldY(13)
	matrix.FoldY(6)
	fmt.Println(matrix.Count())
	matrix.Print()
}
