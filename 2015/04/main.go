package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(md5LeadingZeros(5)) //part one
	fmt.Println(md5LeadingZeros(6)) //part two
}

func md5LeadingZeros(count int) int {
	pad := ""
	for i := 0; i < count; i++ {
		pad += "0"
	}
	input := "iwrupvqb"
	number := 1
	for {
		secret := input + strconv.Itoa(number)
		hashBytes := md5.Sum([]byte(secret))
		hashStr := hex.EncodeToString(hashBytes[:])
		if hashStr[:count] == pad {
			return number
		}
		number++
	}
}
