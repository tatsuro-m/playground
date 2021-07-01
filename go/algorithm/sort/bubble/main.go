package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	s := make([]int, 1000)
	for i := range s {
		s[i] = rand.Intn(100)
	}

	fmt.Println(exec(s))
}

func exec(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}

	return s
}
