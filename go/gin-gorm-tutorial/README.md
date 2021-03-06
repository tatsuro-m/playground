## やったこと

https://qiita.com/Asuforce/items/0bde8cabb30ac094fcb4 を参考にして Gin + GORM で API を作る時のディレクトリ構成を学んだ。

## やりたいこと

- ~~gorm の AutoMigrationではなく https://github.com/golang-migrate/migrate を作ってマイグレートする~~
- ~~環境に応じて接続できる DB を変更できるように環境変数を使う~~
- ~~API のエンドポイントをネストする（/api/v1 みたいな感じで）~~
- ~~id, createdAt などのカラムを追加してみる（gorm.Model は使わなくて良い）~~
- ~~User 構造体に gorm のタグを打って、not null, default値あたりの挙動を確認してみる~~
- ~~gin で json から構造体に bind する時にバリデーションをかけてみる（https://github.com/gin-gonic/gin#model-binding-and-validation）~~
- ~~Post モデルを増やしてみる。 User に従属するモデル~~
- ~~Railsで言うところの `dependent_destroy` 的な挙動はどうやって実現するのかやってみる~~
- ~~テスト専用の DB を使えるようにする~~
- ~~テスト用の fixture ライブラリ、 https://github.com/go-testfixtures/testfixtures を試す（メリットが薄そうだったのでやめた。普通に自作すれば良い）~~
- ~~エンドポイントのテストを書く（golden とかで良いかも）~~
- ~~cognito（jwt）を使って API の認証・認可を実装する~~

## インフラについて
GCP App Engine を試してみた。
基本は Terraform で管理する。とりあえず API をデプロイして動かすところまではできたので良しとするが、以下の改善点がある。
- Cloud SQL がデフォルトのネットワーク上に立っている？ので VPC を作成してそこに立てるというのをやりたい。
- Cloud SQL の SSL が Off になっている。
- Cloud SQL インスタンスにアクセスできるネットワークを `0.0.0.0/0` にしている。
- `gcloud app deploy` は手動で実行する必要がある。
- プロジェクト単位で簡単に削除できるように tfstate を管理しているプロジェクトとリソースをデプロイするプロジェクトを分けている
    - これが正解かはちょっと微妙
- DB migration が Docker コンテナから手動でやっている
    - `postgres://postgres:{password}@{host}/playground-master-db?sslmode=disable` を指定して migration を実行する

最初に terraform apply した時には DB のパブリックIPアドレスは変わることに注意。    
DSN 環境変数の値も変更する必要がある。
