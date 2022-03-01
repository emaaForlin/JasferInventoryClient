FROM golang:1.17.7 AS builder
ARG VERSION=dev
WORKDIR /go/src/app
COPY . .
RUN go build -o main -ldflags=-X=main.version=${VERSION} main.go

FROM debian:stretch-slim
RUN mkdir -p /app/templates
COPY --from=builder /go/src/app/templates /app/templates
COPY --from=builder /go/src/app/config.json /app/config.json
COPY --from=builder /go/src/app/main /app/main
EXPOSE 9091
WORKDIR /app
CMD ["/app/main"]