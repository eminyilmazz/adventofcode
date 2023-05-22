package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(partOne())
}

func partOne() int {
	number := 1
	input := "iwrupvqb"
	for {
		secret := input + strconv.Itoa(number)
		hashBytes := md5.Sum([]byte(secret))
		hashStr := hex.EncodeToString(hashBytes[:])
		if hashStr[:5] == "00000" {
			return number
		}
		number++
	}
}
