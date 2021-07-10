## 目標
- GCP での VPC やサブネットなどの構築を理解する
- GCE エンジンなどをそのネットワークに関連付ける形で立ててみる
- GCE にサービスアカウントを関連付ける
- CloudSQL をプライベートな空間（パブリックIPを持たない）に立ててみる
- GCE から CloudSQL に psql コマンドなどで接続できることを確認する

## 分かったこと
- CloudSQL を private ip でのみアクセス可能にする場合、「プライベート サービス アクセス」というコネクションを利用する
- GCE の os login を有効にするためにはインスタンスに 
```terraform
  metadata = {
    enable-oslogin = "TRUE"
  }
```
の記述が必要。
- GCE のインスタンスタイプは `"e2-micro` がおそらく最安値
- GCE から private な Cloud SQL への接続は問題なく行えた
```shell
$ psql -h {host_name} -U postgres -p 5432 -d postgres
Password for user postgres: 
psql (11.12 (Debian 11.12-0+deb10u1), server 13.2)
WARNING: psql major version 11, server major version 13.
         Some psql features might not work.
SSL connection (protocol: TLSv1.3, cipher: TLS_AES_256_GCM_SHA384, bits: 256, compression: off)
Type "help" for help.
postgres=> 
postgres=> 
```
