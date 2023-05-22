package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	fmt.Println(partOne())
}

func partOne() int {
	count := 0
	lines := readInput()
	for _, line := range lines {
		if hasDoubleLetter(line) && hasThreeVowels(line) && doesNotContainSubstrings(line) {
			count++
		}
	}
	return count
}

func hasDoubleLetter(s string) bool {
	prev := rune(s[0])
	for _, c := range s[1:] {
		if c == prev {
			return true
		}
		prev = c
	}
	return false
}

func hasThreeVowels(s string) bool {
	re := regexp.MustCompile(`[aeiou]`)
	vowels := re.FindAllString(s, -1)
	return len(vowels) >= 3
}

func doesNotContainSubstrings(s string) bool {
	re := regexp.MustCompile(`ab|cd|pq|xy`)
	return !re.MatchString(s)
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
