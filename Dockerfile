# syntax=docker/dockerfile:1

FROM golang:latest as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o go-recipe-api .

EXPOSE 33000

CMD ["./go-recipe-api"]