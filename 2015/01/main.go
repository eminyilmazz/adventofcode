package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}

func partOne() (int, error) {
	file, err := os.Open("input.txt")
	pos := 0
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		for _, char := range data {
			if char == '(' {
				pos++
			} else {
				pos--
			}
		}
	}
	return pos, nil
}

func partTwo() int {
	file, err := os.Open("input.txt")
	pos := 0
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		for i, char := range data {
			if char == '(' {
				pos++
			} else {
				pos--
			}
			if pos == -1 {
				return i + 1
			}
		}
	}
	fmt.Println(pos)
	return -1
}
