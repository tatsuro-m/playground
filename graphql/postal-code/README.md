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

## csv 元データについて。
公式からダウンロードしてきた csv での行数は `124523` 件ある。  
ただし seed を実行してDBに保存されるのは `120383` 件で、４千件ほどのズレがある。  

以下のようなデータもある。
ビルの場合は「○階」などが付与されてそれぞれに郵便番号が割り当てられ場合があるようだが、以下のビルは住所部分が全く同じなのに郵便番号だけが違うというもの。  

```csv
"9806090","宮城県","仙台市　青葉区","中央　ＳＳ３０住友生命仙台中央ビル","MIYAGI KEN","SENDAI SHI AOBA KU","CHUO SS30-SUMITOMOSEIMEISENDAICHUOB"
"9806001","宮城県","仙台市　青葉区","中央　ＳＳ３０住友生命仙台中央ビル","MIYAGI KEN","SENDAI SHI AOBA KU","CHUO SS30-SUMITOMOSEIMEISENDAICHUOB"
"9806002","宮城県","仙台市　青葉区","中央　ＳＳ３０住友生命仙台中央ビル","MIYAGI KEN","SENDAI SHI AOBA KU","CHUO SS30-SUMITOMOSEIMEISENDAICHUOB"
```
