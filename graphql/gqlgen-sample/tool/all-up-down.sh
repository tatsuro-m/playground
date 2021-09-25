#!/bin/zsh

echo '全てのコンテナを down してから up します'
docker compose down
docker compose up
