#!/bin/bash

# ローカル開発用のスクリプト

# 0. マイグレーションファイルの作成
# docker compose run --rm migration migrate create -ext sql -dir migrations -seq migration_file_name
# 作成された sql file に up, down の処理を記述する

# 1. 開発環境 DB に向けて migration up の実行
docker compose run --rm migration

# 2. 変更されたDBスキーマを元に model ファイルを作り直す
docker compose run --rm boiler

# 3. model のテストもあるのでテスト全体を実行
docker compose run --rm test

# 全て通ったら作成されたコードを利用して実装へ
