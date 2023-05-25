package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func lookAndSpeak(input string) string {
	var newStr strings.Builder
	count := 0
	currChar := rune(input[0])

	for _, char := range input {
		if currChar == 0 || currChar == char {
			count++
			continue
		}
		newStr.WriteString(strconv.Itoa(count))
		newStr.WriteRune(currChar)
		count = 1
		currChar = char

	}
	newStr.WriteString(strconv.Itoa(count))
	newStr.WriteRune(currChar)

	return newStr.String()
}

func solution(iteration int) int {
	start := time.Now().UnixMilli()
	input := "1113122113"
	for i := 0; i < iteration; i++ {
		input = lookAndSpeak(input)
	}
	fmt.Println(time.Now().UnixMilli() - start)
	return len(input)
}

func main() {
	fmt.Println(solution(40)) //part one
	fmt.Println(solution(50)) //part two
}
