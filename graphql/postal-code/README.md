## これは何
郵便番号を入力するとその結果が返ってくる GraphQL API を作成する。

## データのセットアップ
1. https://www.post.japanpost.jp/zipcode/dl/roman-zip.html から csv をダウンロードする。
2. seed を動かす（12万件以上あるので完了まで30分以上かかる）
```shell
$ docker compose run --rm seed sh
# go run ./cmd/seed/main.go 
```

## できること
- 郵便番号 -> 住所
- 住所 -> 郵便番号
の取得。細かい部分はゴールデンテスト参照。

## テストに関する注意点
`go test -v ./...` などでまとめて実行した時に DB データの競合が起きて失敗する。  
パッケージごとに並列実行されるのを回避するために、とりあえず `-p 1` オプションを付けること。
