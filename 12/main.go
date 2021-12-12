package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/hallgren/aoc2021"
)

type Type string

const (
	Start Type = "start"
	End   Type = "end"
	Big   Type = "big"
	Small Type = "small"
)

type Node struct {
	Type  Type
	Value string
}

type Result struct {
	Paths [][]Node
}

func CreateNode(s string) Node {
	switch s {
	case "start":
		return Node{Type: Start, Value: s}
	case "end":
		return Node{Type: End, Value: s}
	default:
		if unicode.IsUpper(rune(s[0])) {
			return Node{Type: Big, Value: s}
		} else {
			return Node{Type: Small, Value: s}
		}
	}
}

func CreatedNodes(s string) (Node, Node) {
	p := strings.Split(s, "-")
	n1 := CreateNode(p[0])
	n2 := CreateNode(p[1])
	return n1, n2
}

func SmallVisited(path []Node, v string) bool {
	for _, node := range path {
		if node.Value == v {
			return true
		}
	}
	return false
}

func Run(path []Node, node Node, connections map[string][]Node, res *Result) {
	path = append(path, node)
	if node.Type == End {
		fmt.Println("end", path)
		res.Paths = append(res.Paths, path)
	} else {
		nodeConnections := connections[node.Value]
		for _, conn := range nodeConnections {
			if conn.Type == Big || conn.Type == End {
				Run(path, conn, connections, res)
			} else if conn.Type == Small {
				if !SmallVisited(path, conn.Value) {
					Run(path, conn, connections, res)
				}
			}
		}
	}
}

func main() {
	connections := make(map[string][]Node)
	lines := aoc2021.Lines("input")
	for _, line := range lines {
		node, node2 := CreatedNodes(line)
		connections[node.Value] = append(connections[node.Value], node2)
		connections[node2.Value] = append(connections[node2.Value], node)
	}
	start := Node{Type: Start, Value: "start"}
	res := Result{}
	Run([]Node{}, start, connections, &res)
	fmt.Println(res.Paths, len(res.Paths))
}
