## やったこと

https://qiita.com/Asuforce/items/0bde8cabb30ac094fcb4 を参考にして Gin + GORM で API を作る時のディレクトリ構成を学んだ。

## やりたいこと

- gorm の AutoMigrationではなく https://github.com/golang-migrate/migrate を作ってマイグレートする
- 環境に応じて接続できる DB を変更できるように環境変数を使う
- API のエンドポイントをネストする（/api/v1 みたいな感じで）
- Post モデルを増やしてみる。 User に従属するモデル
- Gorm が自動で作ってくれるカラムを利用する（id, createdAt など）
- テスト専用の DB を使えるようにする
- テスト用の fixture ライブラリ、 https://github.com/go-testfixtures/testfixtures を試す
- エンドポイントのテストを書く（golden とかで良いかも）
- cognito（jwt）を使って API の認証・認可を実装する
- 
