# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /src

COPY go.mod ./

COPY gp.sum ./

RUN go mod download