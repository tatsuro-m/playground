FROM golang:1.17.3-alpine3.15

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
ENV TZ=Asia/Tokyo
WORKDIR ${ROOT}

RUN apk update && apk add git mysql-client
COPY go.mod go.sum ./
RUN go mod download
