# 『Kubernetes完全ガイド（第二版）』 付録マニフェストのリポジトリ

はじめまして。
青山真也（[@amsy810](https://twitter.com/@amsy810)）と申します。

この度は『Kubernetes完全ガイド（ `第二版` ）』を選んで頂き、誠にありがとうございます。
（このリポジトリへのStarもお待ちしております。）

電子書籍に関しては、Amazon以外からもPDFなどで入手可能です。
* Amazon: https://www.amazon.co.jp/dp/4295009792
* Impress: https://book.impress.co.jp/books/1119101148
* その他いくつかのサイト


## （注意）第一版について

こちらのブランチは `第二版` 用です。
`第一版` に関しては、[1st-edition branch](https://github.com/MasayaAoyama/kubernetes-perfect-guide/tree/1st-edition)を参照してください。

```
$ git checkout 1st-edition
```

# Kubernetes perfect guide "2nd edition" - sample manifest repository

Hello, I'm Masaya Aoyama ([@amsy810](https://twitter.com/@amsy810)).

Thank you for choosing "kubernetes perfect guide "2nd edition".
If you like this repo, please add star :)

You can get E-book and PDF from Amazon or various sites.
* Amazon: https://www.amazon.co.jp/dp/4295009792
* Impress: https://book.impress.co.jp/books/1119101148
* etc

## Note: about 1st edition

This branch is for `2nd edition`.
For `1st edition`, please refer [1st-edition branch](https://github.com/MasayaAoyama/kubernetes-perfect-guide/tree/1st-edition).

```
$ git checkout 1st-edition
```

## 環境
本書で使っているのと同じように GKE を利用する。バージョンが１部違うので改めて記述しておく。

### 作成
基本的に結構時間かかる（５分以上？）ので注意。
```shell
$ gcloud container clusters create k8s \
--cluster-version 1.20.8-gke.700 \
--zone asia-northeast1-a \
--num-nodes 3 \
--machine-type e2-micro \
--enable-network-policy \
--enable-vertical-pod-autoscaling
```

### クラスターへの認証情報を取得し直す
これで kubeconfig に認証情報が設定される。
```shell
 $ gcloud container clusters get-credentials k8s --zone asia-northeast1-a
```

### 削除
```shell
 $ gcloud container clusters delete k8s --zone asia-northeast1-a
```
