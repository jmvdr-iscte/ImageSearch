FROM golang:1.21.6

RUN go install github.com/cosmtrek/air@latest
WORKDIR /usr/src/app

COPY . .
RUN go mod tidy