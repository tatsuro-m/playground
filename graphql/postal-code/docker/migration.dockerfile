FROM golang:1.17.3-alpine3.15

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
COPY migrations ${ROOT}/migrations
