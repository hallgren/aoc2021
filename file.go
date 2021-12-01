package aoc2021

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("vim-go")
}

func Lines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
