package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	pos := 0
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
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
	fmt.Println(pos)
}
