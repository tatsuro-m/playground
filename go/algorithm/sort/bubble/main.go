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
	for _, e := range s {
		fmt.Println(e)
	}

	return s
}
