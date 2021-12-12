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

func main() {
	connections := make(map[string][]string)
	lines := aoc2021.Lines("sample")
	for _, line := range lines {
		node, node2 := CreatedNodes(line)
		fmt.Println("node", node, node2)
		connections[node.Value] = append(connections[node.Value], node2.Value)
	}
	fmt.Println("connections", connections)
}
