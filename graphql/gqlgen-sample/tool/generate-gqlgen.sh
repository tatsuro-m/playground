#!/bin/zsh

echo 'gqlgen の generator を起動します'
docker compose run --rm graph_server go run github.com/99designs/gqlgen
