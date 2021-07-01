package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	s := make([]int, 1000)

	for i, _ := range s {
		s[i] = rand.Intn(100)
	}

	fmt.Println(s)
}
