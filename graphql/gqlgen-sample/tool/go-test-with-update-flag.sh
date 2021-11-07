#!/bin/zsh

echo '-update flag 付きでテストを実行する'
docker compose run --rm test go test ./... -update -v
