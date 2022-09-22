FROM golang:1.19.1

ARG HTTP_PROXY=""
ARG HTTPS_PROXY=""
ARG NO_PROXY=""

ENV http_proxy=${HTTP_PROXY}
ENV https_proxy=${HTTPS_PROXY}
ENV no_proxy=${NO_PROXY}

WORKDIR /usr/src/app

COPY src/* ./
RUN go mod download && go mod verify
RUN go build -v -o /usr/local/bin/shmailr ./...
