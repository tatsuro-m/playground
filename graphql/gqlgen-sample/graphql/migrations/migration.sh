#! /bin/sh

mysqldef -h "${DB_HOST}" -u"${DB_USER}" -p"${DB_PASSWORD}" -P "${DB_PORT}" "${DB_NAME}" < ./migrations/schema.sql
