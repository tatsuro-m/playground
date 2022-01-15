#!/bin/zsh

echo 'stg kubernetes initial setup'
cd ~/dev/playground/kubernetes/k8s-playground/infrastructure/kubernetes/
gcloud config set project playground-318023
gcloud container clusters get-credentials stg-k8s-pg-main -z asia-northeast1-a

kubectl apply -k ./argocd/overlays/stg
