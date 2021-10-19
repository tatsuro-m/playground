#! /bin/bash

schemalex -h ${DB_HOST} -u${DB_USER} -p${DB_PASSWORD} -P ${DB_PORT} dev < schema.sql
