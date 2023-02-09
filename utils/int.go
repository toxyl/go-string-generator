package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func GetRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	n := max - min + 1
	if n <= 0 {
		return min
	}
	return rand.Intn(n) + min
}

func IntFromString(str string, def int) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return def
	}
	return i
}
