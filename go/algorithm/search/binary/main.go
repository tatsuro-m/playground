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
	targetVal := rand.Intn(100)
	for i := range ns {
		ns[i] = rand.Intn(100)
	}
	s := uniq(ns)

	// 二分探索法なので、昇順にソートしてからスタート
	// スライスの半分のところで分割して探索していく
	sort.Ints(s)
	fmt.Println(s)
	fmt.Printf("ターゲット値: %d\n", targetVal)

	exec := func([]int) bool {
		var find bool

		//	半分の位置の要素と比較する
		el := s[len(s)/2]
		if targetVal == el {
			fmt.Printf("探していた値は %d が見つかりました。\n", el)
			find = true
			return find
		} else if targetVal < el {
			s = s[:(len(s)/2)-1]
			find = false
			return find
		} else if targetVal > el {
			s = s[len(s)/2:]
			find = false
			return find
		} else {
			find = false
			return find
		}
	}

	for !exec(s) && len(s) != 1 {
		fmt.Println("探索実行")
		fmt.Println(s)
	}
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
