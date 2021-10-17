#!/usr/bin/env zx

await $`echo Dockerイメージをビルドして対応するレジストリに push します！`

const imageTag = Math.floor(new Date().getTime() / 1000)

// common
await $`kubectl apply -f ../infrastructure/kubernets/common -R`

// nginx
let appName = 'nginx'
let registryURI = 'asia-northeast1-docker.pkg.dev/playground-318023/stg-rproxy-nginx-proxy'
await $`cd ../${appName} && docker build -t ${registryURI}/${appName}:${imageTag} . && docker push ${registryURI}/${appName}:${imageTag}`
await $`echo イメージタグを更新します`
await $`cd ../infrastructure/kubernets/${appName}/overlays/stg && kustomize edit set image ${registryURI}/${appName}:${imageTag}`
