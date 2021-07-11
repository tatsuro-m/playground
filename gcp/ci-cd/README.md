## 目標
- Terraform + GCP + GitHub Actions という構成での CI/CD パイプラインの流れを確認する
- Actions で使うサービスアカウントの管理について

期待する動作は以下の通り。
- feature ブランチの push で検証環境に自動デプロイ
- main にマージされたら本番環境に自動デプロイ
- terraform は環境ごとにディレクトリを分ける方法とし、workspace などは利用しない
