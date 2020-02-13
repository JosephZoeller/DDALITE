FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/200106-uta-go/JKJP2/
COPY . .
RUN go get gopkg.in/yaml.v2
RUN go build ./cmd/blandwall/

ENTRYPOINT [ "./blandwall" ]