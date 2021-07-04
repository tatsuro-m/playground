FROM golang:1.16.3-alpine3.13 as dev

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
ENV TZ=Asia/Tokyo
WORKDIR ${ROOT}

RUN apk update && apk add git
EXPOSE 8080

CMD ["go", "run", "main.go"]
