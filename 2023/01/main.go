package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var words = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
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

func partTwo() int {
	lines := readInput("input.txt")

	var result int
	for _, line := range lines {
		numChars := []rune{}
		lineLength := len(line)
		for index, char := range line {
			if unicode.IsDigit(char) {
				numChars = append(numChars, char)
			} else {
				for key, val := range words {
					if index+len(key) <= lineLength && line[index:index+len(key)] == key {
						numChars = append(numChars, val)
					}
				}
			}
		}
		first := numChars[0]
		last := numChars[len(numChars)-1]
		intValue, err := strconv.Atoi(fmt.Sprintf("%c%c", first, last))
		if err != nil {
			log.Fatal(err)
		}
		result += intValue
	}
	return result
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
