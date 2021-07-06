package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	ns := make([]int, 100)
	for i := range ns {
		ns[i] = rand.Intn(100)
	}
	s := uniq(ns)

	sort.Ints(s)
	fmt.Println(s)

	//targetVal := rand.Intn(100)
}

func uniq(ns []int) []int {
	var s []int
	m := make(map[int]bool)
	for _, e := range ns {
		if _, ok := m[e]; !ok {
			m[e] = true
			s = append(s, e)
		}
	}

	return s
}
