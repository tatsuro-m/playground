https://firebase.google.com/docs/auth/admin/verify-id-tokens#web

Firebase Authentication を試す簡単なデモを行ってみる。  
バックエンドにトークンを送信して検証するところまで。

firebase プロジェクトの作成や管理は画面から行った。

---
フロントの方でコンソールに jwt を出すようにしているので、それをバックエンドの `/ping` に送ればトークンを検証してくれてそれに応じたレスポンスを返してくれる。  
デコードに成功した場合には Claims にアクセスしてログに吐き出しているのでそちらも参考にしてください。

一通りの流れは確認できたのでこれでOKとします。
