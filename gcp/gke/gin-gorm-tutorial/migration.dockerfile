FROM golang:1.16.3-alpine3.13

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
COPY migrations ${ROOT}/migrations
