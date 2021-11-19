# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

WORKDIR /app

RUN go get -u github.com/gofiber/cli/fiber

COPY . .

RUN go mod download

CMD [ "fiber", "dev" ]