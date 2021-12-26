package main

import (
	"fmt"
	"strings"

	"github.com/hallgren/aoc2021"
)

type Matrix [101][101][101]bool

type Cube struct {
	X int
	Y int
	Z int
}

type Operation struct {
	On    bool
	Cubes []Cube
}

func Cubes(cordinates []string) []Cube {
	cubes := make([]Cube, 0)
	xMin := aoc2021.Int(cordinates[1])
	xMax := aoc2021.Int(cordinates[2])
	if xMin < -50 || xMin > 50 {
		return []Cube{}
	}
	if xMax < -50 || xMax > 50 {
		return []Cube{}
	}

	yMin := aoc2021.Int(cordinates[4])
	yMax := aoc2021.Int(cordinates[5])
	if yMin < -50 || yMin > 50 {
		return []Cube{}
	}
	if yMax < -50 || yMax > 50 {
		return []Cube{}
	}

	zMin := aoc2021.Int(cordinates[7])
	zMax := aoc2021.Int(cordinates[8])
	if zMin < -50 || zMin > 50 {
		return []Cube{}
	}
	if zMax < -50 || zMax > 50 {
		return []Cube{}
	}

	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			for z := zMin; z <= zMax; z++ {
				cubes = append(cubes, Cube{X: x, Y: y, Z: z})
			}
		}
	}
	return cubes

}

func f(c rune) bool {
	return c == '=' || c == '.' || c == ','
}

func Operations(lines []string) []Operation {
	operations := make([]Operation, 0)
	for _, line := range lines {
		var op bool
		m := strings.Split(line, " ")
		if m[0] == "on" {
			op = true
			fmt.Println("on", op)
		}
		fmt.Printf("Fields are: %q\n", strings.FieldsFunc(m[1], f))
		cubes := Cubes(strings.FieldsFunc(m[1], f))
		if len(cubes) > 0 {
			operations = append(operations, Operation{On: op, Cubes: cubes})
		}

	}
	return operations
}

func main() {
	lines := aoc2021.Lines("input")
	c := Matrix{}
	op := Operations(lines)
	for _, o := range op {
		for _, cube := range o.Cubes {
			c[cube.X+50][cube.Y+50][cube.Z+50] = o.On
		}
	}
	//fmt.Println(c)
	count := 0
	for x := 0; x <= 100; x++ {
		for y := 0; y <= 100; y++ {
			for z := 0; z <= 100; z++ {
				if c[x][y][z] {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
