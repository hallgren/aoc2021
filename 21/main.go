package main

import "fmt"

type Player struct {
	Score    int
	Position int // 0-9 -> 1-10
}

func (p *Player) Add(i int) bool {
	fmt.Println("add", i)
	p.Position += i
	p.Position = p.Position % 10
	fmt.Println("position", p.Position)
	if p.Position == 0 {
		p.Position = 10
	}
	p.Score += p.Position
	fmt.Println("score", p.Score)
	if p.Score >= 1000 {
		return true
	}
	return false
}

type Die struct {
	Value  int
	Rolled int
}

func (d *Die) Roll() int {
	d.Value++
	d.Rolled++
	return d.Value

}

func (d *Die) Roll3() int {
	v := d.Roll() + d.Roll() + d.Roll()
	return v
}

func main() {
	var looserScore int
	p1 := Player{Position: 6}
	p2 := Player{Position: 2}
	die := Die{}
	for {
		v := die.Roll3()
		if p1.Add(v) {
			looserScore = p2.Score
			break
		}
		v = die.Roll3()
		if p2.Add(v) {
			looserScore = p1.Score
			break
		}
	}
	fmt.Println(die.Value, looserScore, die.Value*looserScore)
}
