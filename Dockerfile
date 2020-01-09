FROM golang:alpine AS builder

WORKDIR /go/src/github.com/louis-ver/scorekeep

COPY . .

RUN go build -o /usr/local/bin/scorekeep .

FROM alpine:latest

WORKDIR /usr/local/bin

COPY --from=builder /usr/local/bin/scorekeep scorekeep

ENTRYPOINT [ "scorekeep" ]