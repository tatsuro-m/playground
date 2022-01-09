## これは何
- https://pkg.go.dev/text/template
- https://github.com/spf13/cobra

あたりを利用して大量の GCP サービスアカウントを terraform で作成する CLI を作成する。  
tf module は環境ごとに分ける想定で、 CLI は対話型ではなく環境とマイクロサービス名が書かれた設定ファイルを読み込めば一発で作れるものにする。
