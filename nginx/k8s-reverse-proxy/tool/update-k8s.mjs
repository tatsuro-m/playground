#!/usr/bin/env zx

// イメージの build、push、イメージタグの更新から apply まで全てできるスクリプト
// 完全な作業効率化用かつテスト用なので本番のような環境では絶対に利用しないで下さい

await $`echo Dockerイメージをビルドして対応するレジストリに push します！`

const imageTag = Math.floor(new Date().getTime() / 1000)

// common
await $`kubectl apply -f ../infrastructure/kubernets/common -R`

// nginx
let appName = 'nginx'
let registryURI = 'asia-northeast1-docker.pkg.dev/playground-318023/stg-rproxy-nginx-proxy'
await $`cd ../${appName} && docker build --platform amd64 -t ${registryURI}/${appName}:${imageTag} . && docker push ${registryURI}/${appName}:${imageTag}`
await $`echo イメージタグを更新します`
await $`cd ../infrastructure/kubernets/${appName}/overlays/stg && kustomize edit set image ${registryURI}/${appName}:${imageTag}`
await $`kubectl apply -k ../infrastructure/kubernets/${appName}/overlays/stg`

// frontend
appName = 'frontend'
registryURI = 'asia-northeast1-docker.pkg.dev/playground-318023/stg-rproxy-frontend'
await $`cd ../${appName} && docker build --platform amd64 -t ${registryURI}/${appName}:${imageTag} . && docker push ${registryURI}/${appName}:${imageTag}`
