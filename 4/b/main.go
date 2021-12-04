package main

import (
	"fmt"
	"strings"

	"github.com/hallgren/aoc2021"
)

type Number struct {
	Value  int
	Marked bool
}

type Row struct {
	Numbers []*Number
}

func (r *Row) Winner() bool {
	for _, e := range r.Numbers {
		if !e.Marked {
			return false
		}
	}
	return true
}

func (r *Row) MarkNumber(n int) bool {
	for _, number := range r.Numbers {
		if number.Value == n {
			number.Marked = true
			return true
		}
	}
	return false
}

func (r *Row) AddNumber(n int) {
	number := Number{Value: n}
	r.Numbers = append(r.Numbers, &number)
}

type Board struct {
	Rows []*Row
	Done bool
}

func (b *Board) Winner() bool {
	for _, r := range b.Rows {
		if r.Winner() {
			return true
		}
	}
	return false
}

func (b *Board) MarkNumber(n int) bool {
	for _, r := range b.Rows {
		if r.MarkNumber(n) {
			return true
		}
	}
	return false
}

func (b *Board) AddRow(r *Row) {
	b.Rows = append(b.Rows, r)
}

func (b *Board) Part1(n int) int {
	sum := 0
	for _, r := range b.Rows {
		for _, number := range r.Numbers {
			if !number.Marked {
				sum += number.Value
			}
		}
	}
	return sum * n
}

type Game struct {
	Boards []*Board
	Moves  []int
}

func (g *Game) AllBoardsWinner() bool {
	for _, b := range g.Boards {
		if !b.Done {
			return false
		}
	}
	return true
}

func (g *Game) Play() int {
	for _, move := range g.Moves {
		fmt.Println(move)
		for _, board := range g.Boards {
			if board.MarkNumber(move) {
				if board.Winner() {
					board.Done = true
					if g.AllBoardsWinner() {
						fmt.Println(board.Part1(move))
						return 66
					}
				}
			}
		}
	}
	return 0
}

func Moves(s string) []int {
	r := []int{}
	st := strings.Split(s, ",")
	for _, i := range st {
		r = append(r, aoc2021.Int(i))
	}
	return r
}

func CreateBoard(lines []string) *Board {
	board := Board{}
	for _, line := range lines {
		r := Row{}
		values := strings.Fields(line)
		for _, value := range values {
			v := aoc2021.Int(value)
			r.AddNumber(v)
		}
		board.AddRow(&r)
	}
	return &board
}

func CreateGame(lines [][]string) Game {
	r := []*Board{}
	for _, line := range lines {
		b := CreateBoard(line)
		r = append(r, b)
	}
	g := Game{Boards: r}
	return g
}

func BoardLines(lines []string) [][]string {
	r := [][]string{}
	b := []string{}
	for i, line := range lines {
		if i%6 == 0 {
			// empty line between boards
			b = []string{}
			continue
		}
		b = append(b, line)
		if i%6 == 5 {
			// last line
			r = append(r, b)
		}
	}
	return r
}

func main() {
	lines := aoc2021.Lines("../sample")
	boardLines := BoardLines(lines[1:])

	game := CreateGame(boardLines)
	game.Moves = Moves(lines[0])

	game.Play()
}
