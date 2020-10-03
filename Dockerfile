# Build
FROM golang:1.14-alpine3.11 AS base
COPY . /go/src/go-quiz

WORKDIR /go/src/go-quiz
RUN apk add git gcc

# Go modules
ENV GO111MODULE=on
RUN go mod download

# Compile
RUN go build -a -tags netgo -ldflags '-w' -o /go/bin/go-quiz /go/src/go-quiz/main.go

# Package
FROM alpine:3.11
COPY --from=base /go/bin/go-quiz /go-quiz/go-quiz
COPY --from=base /go/src/go-quiz/migration /go-quiz/migration
WORKDIR /go-quiz
ENTRYPOINT ["/bin/sh"]
