FROM golang:1.20-alpine AS builder

RUN apk --no-cache add bash git make gcc gettext

WORKDIR /usr/local/src

#install deps
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

#build
COPY app ./
RUN go build -o ./bin/app cmd/main.go