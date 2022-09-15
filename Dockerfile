# syntax=docker/dockerfile:1

FROM golang:1.17

WORKDIR /app

COPY src .

COPY go.mod .

COPY go.sum .

RUN go mod download

RUN go build .

RUN ./src

EXPOSE 5000