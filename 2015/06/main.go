package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(partOne())
}

func partOne() int {
	//var grid [1000][1000]bool
	grid := make([][]bool, 1000)
	//idk why I have to do this but I guess that's how you do 2D slice.
	for i := range grid {
		grid[i] = make([]bool, 1000)
	}
	instructions := readInput()
	for _, instruction := range instructions {
		op, start, end := parseInput(instruction)
		for i := start[0]; i <= end[0]; i++ {
			for j := start[1]; j <= end[1]; j++ {
				changeState(&grid, op, i, j)
			}
		}
	}
	return countLights(grid)
}

func changeState(grid *[][]bool, op string, x int, y int) {
	if op == "turn off" {
		(*grid)[x][y] = false
	} else if op == "turn on" {
		(*grid)[x][y] = true
	} else if op == "toggle" {
		(*grid)[x][y] = !(*grid)[x][y]
	}
}

func parseInput(instruction string) (string, [2]int, [2]int) {
	parts := strings.Fields(instruction)
	op := ""
	start := [2]int{}
	end := [2]int{}

	if len(parts) < 4 {
		log.Fatalf("Invalid instruction: %s", instruction)
	}

	switch parts[0] {
	case "turn":
		op = parts[0] + " " + parts[1]
		start = parseCoordinates(parts[2])
		end = parseCoordinates(parts[4])
	case "toggle":
		op = parts[0]
		start = parseCoordinates(parts[1])
		end = parseCoordinates(parts[3])
	default:
		log.Fatalf("Invalid instruction: %s", instruction)
	}

	return op, start, end
}

func parseCoordinates(coordinates string) [2]int {
	parts := strings.Split(coordinates, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("Invalid coordinate: %s", coordinates)
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("Invalid coordinate: %s", coordinates)
	}
	return [2]int{x, y}
}

func countLights(grid [][]bool) int {
	count := 0
	for _, row := range grid {
		for _, isOn := range row {
			if isOn {
				count++
			}
		}
	}
	return count
}

func readInput() []string {
	input, err := os.Open("input.txt")
	var lines []string
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
