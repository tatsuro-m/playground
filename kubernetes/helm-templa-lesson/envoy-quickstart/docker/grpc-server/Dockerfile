FROM golang:1.18.5-alpine AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

COPY ./grpc-server/go.mod ./go.mod
COPY ./grpc-server/go.sum ./go.sum
RUN go mod download

COPY ./grpc-server/ .
RUN go build -o /app/grpc-server /app/cmd/server


FROM alpine:latest AS tool

ENV GRPC_HEALTH_PROBE_VERSION=v0.3.2
RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
  chmod +x /bin/grpc_health_probe


FROM alpine:3.12
RUN apk add --no-cache tzdata

COPY --from=builder /app/grpc-server /app/grpc-server
COPY --from=tool /bin/grpc_health_probe /bin/grpc_health_probe

CMD ["/app/grpc-server"]
