FROM golang:1.16.7-alpine3.14

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
COPY internal/pkg/db/migrations/mysql ${ROOT}/mysql
