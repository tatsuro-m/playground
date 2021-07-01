package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestExec(t *testing.T) {
	t.Run("昇順にソートされていること", func(t *testing.T) {
		rand.Seed(time.Now().Unix())
		s := make([]int, 1000)
		for i := range s {
			s[i] = rand.Intn(100)
		}

		sort.Ints(s)
		want := exec(s)

		if !reflect.DeepEqual(s, want) {
			t.Errorf("結果が違います。")
		}
	})
}
