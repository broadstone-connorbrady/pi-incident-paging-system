FROM golang:1.17-buster AS builder

ENV GOARCH arm
ENV GOOS linux

RUN apt-get install git wget

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o pigs

FROM arm64v8/debian:buster-slim AS image

WORKDIR /app

COPY --from=builder /app/pigs .
COPY pocsag pocsag
RUN chmod +x pocsag