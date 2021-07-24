https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app?hl=ja

をやる。

ただし CloudShell ではなくローカル環境で、かつ terraform で構築するようにする。

---
Dockerfile は古いからか上手くいかなかったので適当に修正した。
バイナリにして実行されるのが確認できたのでそれでOK。

build は以下のコマンドで。
```shell
docker build -t asia-northeast1-docker.pkg.dev/playground-318023/stg-gke-basic-tutorial-hello-repo/hello-app:v1 . --platform linux/amd64
```

