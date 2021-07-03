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
	// スライスの約半分のところで分割する
	left := asc(s[:mid])
	right := asc(s[mid:])
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		// 初回は、２つに分割したスライスそれぞれの先頭の要素を比べる（i, j 共に０で初期化してあるから）
		// もし分割した左側の方が大きいなら、とりあえず右側に
		if left[i] > right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
