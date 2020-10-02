FROM golang:1.15.2-alpine3.12

WORKDIR /app

ENTRYPOINT ["go", "run", "main.go"]