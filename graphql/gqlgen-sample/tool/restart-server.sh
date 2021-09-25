#!/bin/bash

echo "gin と graphql サーバーの両方を再起動します"
docker compose restart gin_web
docker compose restart graph_server
