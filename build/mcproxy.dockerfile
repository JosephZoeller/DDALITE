FROM golang:rc-buster AS builder

WORKDIR $GOPATH/src/github.com/200106-uta-go/JKJP2
COPY . . 
RUN pwd
RUN ls ./pkg/proxy

RUN go build ./cmd/mcproxy

FROM debian:latest
RUN mkdir /home/vnf
COPY --from=builder $GOPATH/src/github.com/200106-uta-go/JKJP2 /
RUN  mv /home/vnf/mcproxy /bin/
EXPOSE 80 443 22
EXPOSE 4444
RUN cd /home/vnf/