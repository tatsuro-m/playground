FROM golang:1.18.4 AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

COPY ./api/go.mod ./
RUN go mod download

COPY ./api .
RUN go build -o /app/api

FROM alpine:3.15.5

RUN apk add --no-cache tzdata
COPY --from=builder /app/api /app/api

CMD ["/app/api"]
