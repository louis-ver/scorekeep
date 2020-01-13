FROM golang:alpine AS builder

WORKDIR /go/src/github.com/louis-ver/scorekeep

COPY . .

RUN apk add git \
    && go get -u -v \
        github.com/gin-gonic/gin \
        github.com/jinzhu/gorm \
        github.com/jinzhu/gorm/dialects/postgres \
    && go build -o /usr/local/bin/scorekeep .

FROM alpine:latest

WORKDIR /usr/local/bin

COPY --from=builder /usr/local/bin/scorekeep scorekeep

ENTRYPOINT [ "scorekeep" ]