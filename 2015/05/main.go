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
	fmt.Println(partTwo())
	fmt.Println(hasPairRepeating("ueihvxviirnooomi"))
}

func partTwo() int {
	count := 0
	lines := readInput()
	for _, line := range lines {
		if hasRepeatingWithOneBetween(line) && hasPairRepeating(line) {
			count++
		}
	}
	return count
}

func hasRepeatingWithOneBetween(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

// I should have used regexp
// but regexp depresses me
// instead I preferred to get lost in substrings
func hasPairRepeating(s string) bool {
	pairs := make(map[string]bool)
	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		//check for overlap
		//skip if it overlaps
		if i+2 < len(s) && pair == s[i+1:i+3] {
			continue
		}
		if pairs[pair] {
			return true
		}
		pairs[pair] = true
	}
	return false
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
