FROM golang:1.17.1-buster

RUN apt install make git gcc

LABEL REPO="https://github.com/lowellmower/audmon"

ENV PROJPATH=/go/src/github.com/lowellmower/audmon
# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

COPY . /go/src/github.com/lowellmower/audmon
WORKDIR /go/src/github.com/lowellmower/audmon
VOLUME /go/src/github.com/lowellmower/audmon/bin

ENTRYPOINT make build