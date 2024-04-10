package utils

import (
	"math/rand/v2"
	"strconv"
)

func GetRandomInt(min, max int) int {
	n := max - min + 1
	if n <= 0 {
		return min
	}
	return rand.IntN(n) + min
}

func IntFromString(str string, def int) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return def
	}
	return i
}
