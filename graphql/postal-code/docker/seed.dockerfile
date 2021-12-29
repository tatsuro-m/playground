FROM golang:1.17.3-alpine3.15 as dev

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
ENV TZ=Asia/Tokyo
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download
EXPOSE 8080

CMD ["go", "run", "./cmd/seed/main.go"]
