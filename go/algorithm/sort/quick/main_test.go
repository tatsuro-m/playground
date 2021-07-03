package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestAsc(t *testing.T) {
	t.Run("昇順にソートされていること", func(t *testing.T) {
		rand.Seed(time.Now().Unix())
		s := make([]int, 30)
		for i := range s {
			s[i] = rand.Intn(100)
		}
		c := make([]int, 30)
		copy(c, s)

		asc(s)
		sort.Ints(c)

		if !reflect.DeepEqual(c, s) {
			t.Errorf("結果が違います。\n actual: %v \n want: %v", s, c)
		}
	})
}
