# build stage
FROM golang:latest

WORKDIR /app

ENV GO111MODULE=on

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o hypertyperbot

ENTRYPOINT ["/app/hypertyperbot"]
