FROM golang:latest

RUN go build -buildvcs=false -o ./app

ENTRYPOINT ./app