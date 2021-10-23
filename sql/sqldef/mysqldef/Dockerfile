FROM alpine:3.14

WORKDIR /usr/local/bin
RUN apk add mysql-client && \
  wget https://github.com/k0kubun/sqldef/releases/download/v0.11.10/mysqldef_linux_amd64.tar.gz -O mysqldef.tar.gz && \
  tar -xzvf mysqldef.tar.gz

COPY ./schema.sql ./migration.sh ./
