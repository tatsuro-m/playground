FROM golang:1.17.0-alpine3.14

ENV ROOT=/go/src/app
ENV TZ=Asia/Tokyo
WORKDIR ${ROOT}

RUN apk update && apk add git postgresql-client

RUN go install github.com/volatiletech/sqlboiler/v4@latest && \
    go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

CMD ["sqlboiler", "psql"]
