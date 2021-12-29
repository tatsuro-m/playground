FROM golang:1.17.3-alpine3.15

ENV ROOT=/go/src/app
ENV TZ=Asia/Tokyo
WORKDIR ${ROOT}

RUN apk update && apk add git mysql-client

RUN go install github.com/volatiletech/sqlboiler/v4@latest && \
    go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest

CMD ["sqlboiler", "mysql"]
