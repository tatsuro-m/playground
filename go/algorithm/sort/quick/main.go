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
	// 単純に先頭、中間、最後の要素の中から中間の値を地道な if で洗い出しているだけ
	pivot := Med(s[0], s[len(s)/2], s[len(s)-1])

	//　left, right 共にインデックス番号として使う
	left := 0
	right := len(s) - 1

	// 条件式を書かないことによって無限ループにする
	for {
		//　条件式だけ書いているのでいわゆる while 文になる
		// 左から探索していって、値が基準値よりも小さい限りは１個ずつ対象を右にずらしていく。
		for s[left] < pivot {
			left++
		}
		//　ループを抜けているので、この時点で left が表すのは、「先頭から探索して最初に見つかった、基準値以上の要素のインデックス番号」になっているはず。

		// 右から探索していって、値が基準値よりも大きい限りは１個ずつ対象を左にずらしていく。
		for s[right] > pivot {
			right--
		}
		// ループを抜けているので、この時点で right が表すのは、「末尾から探索して最初に見つかった、基準値以下の要素のインデックス番号」になっているはず。

		//	左右からの探索が交差したら終了するが、交差していない場合にはこの後の処理が続く
		if left >= right {
			break
		}

		// ここまで来ているので、s[left]　の値は基準値以上である。同様に s[right] の値は基準値以下である。
		// 昇順にソートしたいので、 s[left] が後ろに来なくてはいけないので交換する。
		s[left], s[right] = s[right], s[left]

		flag := true
		if s[right] == pivot {
			left++
			flag = false
		}
		if s[left] == pivot && flag {
			right--
		}
	}

	s1 := s[:left]
	if len(s1) > 1 {
		asc(s1)
	}

	s2 := s[right+1:]
	if len(s2) > 1 {
		asc(s2)
	}

	cnt := 0
	for _, v := range s1 {
		s[cnt] = v
		cnt++
	}
	s[cnt] = pivot
	cnt++

	for _, v := range s2 {
		s[cnt] = v
		cnt++
	}

	return s
}

func Med(x, y, z int) int {
	if x < y {
		if y < z {
			return y
		} else if x < z {
			return z
		} else {
			return x
		}
	} else {
		if x < z {
			return x
		} else if y < z {
			return z
		} else {
			return y
		}
	}
}
