# 実際の DB を使って行うテストでデータをどのようにクリーンにするか問題
調べてみる。
https://qiita.com/seya/items/582c2bdcca4ad50b03b7#3%E3%83%86%E3%82%B9%E3%83%88%E3%83%87%E3%83%BC%E3%82%BF%E3%81%AE%E3%82%AF%E3%83%AA%E3%83%BC%E3%83%B3%E3%83%8A%E3%83%83%E3%83%97%E3%82%92%E3%81%A9%E3%81%86%E3%82%84%E3%81%A3%E3%81%A6%E8%A1%8C%E3%81%86%E3%81%8B

この記事は非常に参考になる。  
主には以下２つの方針がある。
1. テストケースごとに全てのテーブルを TRUNCATE する
2. DB をテストケースごとに CREATE, DROP する。もちろんマイグレーションも。

どちらでも良さそうだが、速度の面で比較してみたい。  

## 検証
`go/gin-gorm-tutorial/controller/user/user_test.go` のテストファイルを使う。
```go
	for i := 0; i < 500; i++ {
		for _, tt := range tests {
```
結構適当だけど、500 * テストパターンの4つで 2000回ぐらい実行されるはず。
この同じ条件で比較する。
`docker compose exec` を使って実行する。最初の実行と２回目をすぐに実行した時の２つで比較する。

### 1のパターン
1回目: 
```shell
real    2m58.053s
user    0m0.111s
sys     0m0.055s

Process finished with exit code 0
```
2回目
```shell
real    0m0.822s
user    0m0.103s
sys     0m0.052s

Process finished with exit code 0
```
２回目爆速すぎるんだけど、ちゃんと実行されてる？と心配になるレベル。こんなもん？

### 2のパターン
意外と面倒で最後までいきませんでした。  
が、ループ回数を１回にしてやった時点でこれですから、おそらくこちらの方が遅いのでは無いかと。
```shell
real    0m0.841s
user    0m0.097s
sys     0m0.039s

Process finished with exit code 1
```
CIのことも考えると、その都度マイグレーションまで実行するのは微妙かも。  
まだまだ改善の余地はあるけど、とりあえず１の方法でやっていくことにしましょう。
