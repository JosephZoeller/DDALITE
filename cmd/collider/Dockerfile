FROM golang:alpine AS builder


# run dockerfile from root directory
WORKDIR $GOPATH/github.com/200106-uta-go/JKJP2/
COPY . .
# TIL $GOPATH refers to the docker machine's gopath, even when using it with the source path

RUN apk add --no-cache git
RUN go get -d
RUN go build -o Collider ./cmd/collider

ADD ./dictionary.txt ./dictionary.txt

FROM alpine:latest

COPY --from=builder dictionary.txt .
COPY --from=builder Collider .
CMD [ "./Collider" ]