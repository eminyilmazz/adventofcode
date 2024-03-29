package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	R = 12
	G = 13
	B = 14
)

var (
	minR int
	minG int
	minB int
)

type game struct {
	r int
	g int
	b int
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}

func partOne() int {
	var val int
	inputs := readInput()
	var possible bool

	for i, input := range inputs {
		possible = false
		line := getGames(input)
		games := strings.Split(line, ";")
		for _, game := range games {
			possible = true
			if !isPossible(game) {
				possible = false
				break
			}
		}
		if possible {
			val = val + i + 1
		}
	}
	return val
}

func partTwo() int {
	var val int
	inputs := readInput()

	for _, input := range inputs {
		line := getGames(input)
		games := strings.Split(line, ";")
		for _, game := range games {
			getMinRequired(game)
		}
		val = val + (minR * minG * minB)
		minR = 0
		minG = 0
		minB = 0
	}
	return val
}

func getGames(l string) string {
	games := strings.Split(l, ":")[1]
	return games
}

func isPossible(s string) bool {
	cubes := strings.Split(s, ",")
	var g game

	for _, cube := range cubes {
		count, color := parseColor(cube)
		switch color {
		case "red":
			g.r += count
		case "green":
			g.g += count
		case "blue":
			g.b += count
		}
	}
	return g.r <= R && g.g <= G && g.b <= B
}

func parseColor(cube string) (int, string) {
	parts := strings.Split(strings.TrimSpace(cube), " ")
	count, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	return count, parts[1]
}

func getMinRequired(s string) {
	cubes := strings.Split(s, ",")

	for _, cube := range cubes {
		count, color := parseColor(cube)
		switch color {
		case "red":
			if minR < count {
				minR = count
			}
		case "green":
			if minG < count {
				minG = count
			}
		case "blue":
			if minB < count {
				minB = count
			}
		}
	}
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
