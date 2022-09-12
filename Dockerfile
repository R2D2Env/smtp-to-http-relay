FROM golang:1.19.1

WORKDIR /usr/src/app

COPY src/* ./
RUN go mod download && go mod verify
RUN go build -v -o /usr/local/bin/shmailr ./...
