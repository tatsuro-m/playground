GraphQL Server (gqlgen) で orm も利用したバックエンドを書いてみる。
ORM はとりあえず sqlboiler でやる。
resolver からは他のマイクロサービスを叩くのではなく直接 RDB に接続する構成にする。  
余裕があれば簡単なフロントエンドアプリも。

認証には jwt トークンを利用する。
