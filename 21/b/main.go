package main

import (
	"fmt"
	"time"
)

type Player struct {
	Score    int
	Position int // 0-9 -> 1-10
}

func (p *Player) Add(i int) bool {
	p.Position += i
	p.Position = p.Position % 10
	if p.Position == 0 {
		p.Position = 10
	}
	p.Score += p.Position
	if p.Score >= 21 {
		return true
	}
	return false
}

type Score struct {
	P1 int
	P2 int
}

type Game struct {
	P1, P2 Player
	next   int
	Score  *Score
}

func Roll(g Game, count int) {
	// start universes
	m := make(map[int]int)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for z := 1; z <= 3; z++ {
				m[i+j+z]++
			}
		}

	}
	for k, v := range m {
		Play(g, k, v*count)
	}
}

func Play(g Game, v, count int) {
	if g.next%2 == 0 {
		if g.P1.Add(v) {
			// set winner
			g.Score.P1 += count
			//fmt.Println("P1 winner", g.P1.Score)
			return
		}
	} else {
		if g.P2.Add(v) {
			// set winner
			g.Score.P2 += count
			//fmt.Println("P2 winner", g.P2.Score)
			return
		}
	}
	g.next++
	Roll(g, count)
}

func main() {
	p1 := Player{Position: 6}
	p2 := Player{Position: 2}
	score := Score{}
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(score.P1, score.P2)
		}
	}()
	game := Game{P1: p1, P2: p2, Score: &score}
	Roll(game, 1)
	fmt.Println(score.P1, score.P2)
}
