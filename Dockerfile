FROM golang:alpine AS builder

WORKDIR /go/src/github.com/louis-ver/scorekeep

COPY . .

ENV TZ="America/Montreal"

RUN apk add \
    git \
    tzdata \
    && go get -u -v \
        github.com/gin-gonic/gin \
        github.com/gocolly/colly \
        golang.org/x/text/transform \
        golang.org/x/text/unicode/norm \
    && go build -o /usr/local/bin/scorekeep-server ./scorekeep-server

FROM alpine:latest

WORKDIR /usr/local/bin

COPY --from=builder /usr/local/bin/scorekeep-server scorekeep-server

ENTRYPOINT [ "scorekeep-server" ]