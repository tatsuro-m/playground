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
	//　要素数が１以下なら何もせずスライスを返す
	if len(s) <= 1 {
		return s
	} else {
		// pivot の求め方は今回適当で、単純に先頭の要素を採用する。
		pivot := s[0]

		place := 0

		// s のインデックス番号の末尾数まで繰り返す
		for j := 0; j < len(s)-1; j++ {
			// 先頭要素は pivot として使っているので、今回は２番目の要素から探索する
			if s[j+1] < pivot {
				// もし基準値の方が大きいのなら入れ替える。
				// place は０で初期化してあって、入れ替えるとインクリメントされる。
				s[j+1], s[place+1] = s[place+1], s[j+1]
				place++
			}
		}
		s[0], s[place] = s[place], s[0]

		first := asc(s[:place])
		second := asc(s[place+1:])
		first = append(first, s[place])

		first = append(first, second...)
		return first
	}
}
