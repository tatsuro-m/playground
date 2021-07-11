## 目標
- 手を動かして IAM を理解する
  - サービスアカウントの作成
  - メンバーとロールの binding 方法
  - あるエンティティにどの権限が付与されているかを確認する方法 
- 組織、ディレクトリを使った権限管理モデルについて理解する
- 権限の継承について

## 分かったこと
- terraform の iam 関連設定は色々あるが、「誰が、どのアクションを、何に対して」行えるのかということを意識すると理解しやすい
- リソースとして見るサービスアカウントに対してアイデンティティとして見るサービスアカウントからのアクセス権限を付与するみたいなこともある
- 基本的には project 全体に対して特定のロールを許可するアクションをメンバーに付与するのが良い
  - https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/google_project_iam
  - 個々のリソースレベルまで絞り込みたいなら条件式を使うか、サービスごとに用意されているサービスアカウントを関連付けるリソースを利用する
- 「Aというサービスアカウントへのアクセス権限があるメンバーを調べたい」ということならコンソールから簡単に分かるが、「A に紐付いているポリシーを確認したい」という場合には cli を使う必要があるかも
  - https://cloud.google.com/asset-inventory/docs/searching-iam-policies
```shell
$ gcloud asset search-all-iam-policies --scope='projects/iam-lesson-6359'  --query='policy : "serviceAccount:c-967479@iam-lesson-6359.iam.gserviceaccount.com"'
---
assetType: cloudresourcemanager.googleapis.com/Project
policy:
  bindings:
  - members:
    - serviceAccount:c-967479@iam-lesson-6359.iam.gserviceaccount.com
    role: roles/run.admin
  - members:
    - serviceAccount:c-967479@iam-lesson-6359.iam.gserviceaccount.com
    role: roles/storage.admin
project: projects/109418192817
resource: //cloudresourcemanager.googleapis.com/projects/iam-lesson-6359
```
- `Authoritative` な記述に注意する。既存のものを削除する動きをするから。回避したい場合には、 `google_*_iam_member` の方を利用すると良い。
- 組織を使うと権限管理の面でメリットが大きいが、個人アカウント（gmail.com）ではそもそも利用できない
  - https://cloud.google.com/resource-manager/docs/creating-managing-organization?hl=JA
  - Google Workspace と関連付けるなどして利用するものが組織
- フォルダ機能も組織があって初めて使えるもののようなので、個人の場合は「組織なし」の直下にプロジェクトを生やすことになる
- 権限の継承は組織が無いので試せていませんが、特に難しくなさそう。特徴的なのは上で指定した権限を下で剥奪したりできないこと。上には最低限共通で必要になる権限を付与するのが良いのかなと。
  
