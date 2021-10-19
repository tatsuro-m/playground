#! /bin/sh

cat ./schema.sql | \
 ./schemalex_linux_amd64/schemalex "mysql://${DB_USER}:${DB_PASSWORD}@tcp(db:3306)/dev" - > ./result_schema.sql && \
 mysql -h ${DB_HOST} -u${DB_USER} -p${DB_PASSWORD} -P ${DB_PORT} dev < ./result_schema.sql
