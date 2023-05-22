package main

import (
	"fmt"
	"os"
)

type House struct {
	x, y int
}

func main() {
	fmt.Println(travel())
}

func readInput() string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	return string(content)
}

func travel() int {
	been := make(map[House]bool)
	directions := readInput()
	current := House{x: 0, y: 0}
	been[current] = true
	for _, d := range directions {
		visit(&been, d, &current)
	}
	return len(been)
}

func visit(been *map[House]bool, d rune, current *House) {
	switch d {
	case '^':
		current.y++ // Move north
	case 'v':
		current.y-- // Move south
	case '>':
		current.x++ // Move east
	case '<':
		current.x-- // Move west
	}
	(*been)[*current] = true
}
