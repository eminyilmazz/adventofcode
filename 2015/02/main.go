package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dimensions := parseLine(line)
		sum += calculateTotal(dimensions)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
	fmt.Println(sum)
}

func parseLine(line string) []int {
	dimensionStrings := strings.Split(line, "x")
	dimensionSlice := make([]int, len(dimensionStrings))

	for i, dimStr := range dimensionStrings {
		dim, err := strconv.Atoi(dimStr)
		if err != nil {
			fmt.Printf("Error converting dimension %s to integer: %v\n", dimStr, err)
			return nil
		}
		dimensionSlice[i] = dim
	}
	return dimensionSlice
}

func calculateTotal(d []int) int {
	surfaces := getSurfaces(d)
	slack := findSmallest(surfaces)

	sum := 0
	for _, num := range surfaces {
		sum += num
	}
	return sum + slack
}

func getSurfaces(d []int) []int {
	surfaces := make([]int, 3)
	surfaces[0] = 2 * d[0] * d[1]
	surfaces[1] = 2 * d[1] * d[2]
	surfaces[2] = 2 * d[2] * d[0]
	return surfaces
}

func findSmallest(s []int) int {
	smallest := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] < smallest {
			smallest = s[i]
		}
	}
	return smallest / 2
}
