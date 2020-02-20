FROM golang:alpine AS builder


# run dockerfile from root directory
RUN mkdir -p $GOPATH/src/github.com/200106-uta-go/JKJ2/
WORKDIR $GOPATH/src/github.com/200106-uta-go/JKJP2/
COPY . .
# TIL $GOPATH refers to the docker machine's gopath, even when using it with the source path

RUN apk add --no-cache git
RUN pwd
RUN ls
RUN go get -u 'github.com/JosephZoeller/cityhash'
RUN go get -u 'github.com/lib/pq'
RUN go get -u 'gopkg.in/yaml.v3'

RUN go build -o Collider ./cmd/collider

ADD ./dictionary.txt ./dictionary.txt

FROM alpine:latest

COPY --from=builder dictionary.txt .
COPY --from=builder Collider .
CMD [ "./Collider" ]
