package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(partOne())
}

func partOne() int {
	inputs := readInput("input.txt")
	var sum int
	for _, input := range inputs {
		digits := extractDigits(input)
		sum += extractValue(digits)
	}
	return sum
}

func readInput(dest string) []string {
	input, err := os.Open(dest)
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

func extractDigits(input string) string {
	var digits string

	for _, char := range input {
		if unicode.IsDigit(char) {
			digits = digits + string(char)
		}
	}
	return digits
}

func extractValue(digits string) int {
	firstDigit := string([]rune(digits)[0])
	lastDigit := string([]rune(digits)[len(digits)-1])

	first, _ := strconv.Atoi(firstDigit)
	end, _ := strconv.Atoi(lastDigit)
	return first*10 + end
}
