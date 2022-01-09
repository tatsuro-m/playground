#!/bin/sh

# docker compose migration サービスの起動用スクリプト
migrate -database mysql://root:password@tcp\(db:3306\)/dev -path migrations/ up && \
 migrate -database mysql://root:password@tcp\(db:3306\)/test -path migrations/ up

## down バージョン
#migrate -database mysql://root:password@tcp\(db:3306\)/dev -path migrations/ down && \
# migrate -database mysql://root:password@tcp\(db:3306\)/test -path migrations/ down
