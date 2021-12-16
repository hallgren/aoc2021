package main

import (
	"fmt"

	"github.com/hallgren/aoc2021"
)

const Size = 100

type Entry struct {
	Risk       int
	SumRisk    int
	Neighbours []*Entry
}

type Unvisited struct {
	List []*Entry
}

type Matrix [Size][Size]*Entry

func (m Matrix) MaxX() int {
	return Size - 1
}

func (m Matrix) MaxY() int {
	return Size - 1
}

func (m Matrix) Neighbours(y, x int) []*Entry {
	res := []*Entry{}
	// left
	if x-1 >= 0 {
		res = append(res, m[y][x-1])
	}
	// under
	if y+1 <= m.MaxY() {
		res = append(res, m[y+1][x])
	}
	// above
	if y-1 >= 0 {
		res = append(res, m[y-1][x])
	}
	// right
	if x+1 <= m.MaxX() {
		res = append(res, m[y][x+1])
	}
	return res
}

func Create(lines []string) (Matrix, *Unvisited) {
	m := Matrix{}
	// base matrix
	for y, line := range lines {
		for x, value := range line {
			m[y][x] = &Entry{Risk: aoc2021.Int(string(value)), SumRisk: 1000000, Neighbours: []*Entry{}}
		}
	}
	m[0][0] = &Entry{Risk: 0, SumRisk: 0}

	// neighbours and unvisited list
	for y := 0; y < Size; y++ {
		for x, e := range m[y] {
			e.Neighbours = m.Neighbours(y, x)
		}
	}
	return m, m.Unvisited()
}

func (m Matrix) Unvisited() *Unvisited {
	res := []*Entry{}
	for y := 0; y <= m.MaxY(); y++ {
		for x := 0; x <= m.MaxX(); x++ {
			res = append(res, m[y][x])
		}
	}
	return &Unvisited{List: res}
}

func Lowest(unvisited *Unvisited) *Entry {
	var entry *Entry
	var index int
	for i, e := range unvisited.List {
		if entry == nil {
			entry = e
			index = i
		} else if entry.SumRisk > e.SumRisk {
			entry = e
			index = i
		}
	}
	// Remove the element at index i from a.
	unvisited.List[index] = unvisited.List[len(unvisited.List)-1] // Copy last element to index i.
	unvisited.List[len(unvisited.List)-1] = nil                   // Erase last element (write zero value).
	unvisited.List = unvisited.List[:len(unvisited.List)-1]       // Truncate slice.
	return entry
}

func Run(unvisited *Unvisited) {
	for {
		if len(unvisited.List) == 0 {
			break
		}
		lowest := Lowest(unvisited)
		for _, n := range lowest.Neighbours {
			if lowest.SumRisk+n.Risk < n.SumRisk {
				n.SumRisk = lowest.SumRisk + n.Risk
			}
		}
	}
}

func main() {
	lines := aoc2021.Lines("input")
	m, unvisited := Create(lines)
	//fmt.Println("matrix", m, m.MaxX(), m.MaxY(), unvisited)
	Run(unvisited)
	fmt.Println("end", m[m.MaxY()][m.MaxX()], m.MaxX(), m.MaxY())
}
