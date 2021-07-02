package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	s := make([]int, 100)
	for i := range s {
		s[i] = rand.Intn(100)
	}

	fmt.Println(asc(s))
	fmt.Println(desc(s))
}

func asc(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			// 先頭と２つ目（i + 1 だから）を比較して、１つ目の方が大きいかをチェック
			//　昇順したいので、大きい要素は小さい要素よりも後ろに来なくてはいけない
			if s[i] > s[j] {
				// 単純に前後を入れ替えているだけ
				// slice に対してインデックス番号を指定して書き込んでいるので slice 自体の順番が整っていく
				s[i], s[j] = s[j], s[i]
			}
		}
		// 以降 i をインクリメントしながらスライスの長さ分だけ繰り返すので、比較する要素が（１個目と２個目、２個目と３個目 ...）のように変わっていく
	}

	return s
}

func desc(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] < s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}

	return s
}
