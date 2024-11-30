FROM golang:1.23.3-bullseye

RUN apt-get update && apt-get install -y protobuf-compiler

WORKDIR /go/src/go-auth