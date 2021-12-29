FROM golang:1.17.3-alpine3.15 as dev

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
ENV TZ=Asia/Tokyo
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download
EXPOSE 8080

CMD ["go", "run", "./cmd/web/main.go"]

FROM golang:1.17.3-alpine3.15 as builder

ENV ROOT=/go/src/app
ENV TZ=Asia/Tokyo
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

COPY . ${ROOT}
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $ROOT/binary $ROOT/cmd/web


FROM alpine:3.15.0 as prod

ENV ROOT=/go/src/app
ENV TZ=Asia/Tokyo
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

EXPOSE 8080
CMD ["/go/src/app/binary"]
