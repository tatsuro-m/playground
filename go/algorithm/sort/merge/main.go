package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	s := make([]int, 30)
	for i := range s {
		s[i] = rand.Intn(100)
	}

	asc(s)
}

func asc(s []int) []int {
	var result []int
	if len(s) < 2 {
		return s
	}

	mid := len(s) / 2
	r := asc(s[:mid])
	l := asc(s[mid:])
	i, j := 0, 0

	for i < len(r) && j < len(l) {
		if r[i] > l[j] {
			result = append(result, l[j])
			j++
		} else {
			result = append(result, r[i])
			i++
		}
	}

	result = append(result, r[i:]...)
	result = append(result, l[j:]...)

	return result
}
