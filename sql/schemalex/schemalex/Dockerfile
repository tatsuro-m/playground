FROM alpine:3.14

WORKDIR /usr/local/bin
RUN apk add mysql-client && \
  wget https://github.com/schemalex/schemalex/releases/download/v0.1.1/schemalex_linux_amd64.tar.gz -O schemalex.tar.gz && \
  tar -xzvf schemalex.tar.gz

COPY ./schema.sql ./migration.sh ./
