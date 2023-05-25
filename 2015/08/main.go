package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(partOne())
}

func partOne() int {
	lines := readInput()
	var total, actual int
	for _, line := range lines {
		actual = len(line) - 2
		actual = actual - checkEscape(line)
		total = total + len(line) - actual
	}
	return total
}

func checkEscape(line string) int {
	var count int
	for i := 0; i < len(line); i++ {
		charAt := line[i]
		if charAt == '\\' {
			if line[i+1] == 'x' {
				i += 3
				count += 3
			} else {
				count++
				i++
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
