package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	ns := make([]int, 10)
	for i := range ns {
		ns[i] = rand.Intn(10)
	}
	targetVal := rand.Intn(10)

	// 一直線（ linear ）に探していくアルゴリズム
	//	非常にシンプルだが効率はあまり良くない
	for _, e := range ns {
		if e == targetVal {
			fmt.Printf("対象の数字が見つかりました！\n要素: %d\nターゲット: %d\n", e, targetVal)
			return
		}
	}

	fmt.Println("最後まで見つかりませんでした")
}
