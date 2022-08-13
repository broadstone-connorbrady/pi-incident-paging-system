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

RUN apt-get update && apt-get install -y git

RUN apt-get install -y sudo

WORKDIR rpitx
RUN git clone https://github.com/F5OEO/rpitx .
RUN bash install.sh

COPY --from=builder /app/pigs .
