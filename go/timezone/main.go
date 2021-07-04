package main

import (
	"fmt"
	"time"
)

//　結論、環境変数 TZ を "Asia/Tokyo" で設定しておけばOK。
func main() {
	t := time.Now()
	// Mac で直接動かすと、2021-07-04 09:38:28.632678 +0900 JST m=+0.000086751　のようになる。
	// Docker で動かすと（tz を指定しないと） 2021-07-04 00:42:08.726433637 +0000 UTC m=+0.000037834 になる。
	fmt.Println(t)

	loc, _ := time.LoadLocation("Asia/Tokyo")
	// tz を指定していない Docker 環境でもこれなら jst になる。
	fmt.Println(time.Now().In(loc))
}
