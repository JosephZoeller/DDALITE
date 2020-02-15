FROM golang:rc-buster AS builder

RUN mkdir -p /home/go/src/github.com/200106-uta-go/JKJP2
WORKDIR /home/go/src/github.com/200106-uta-go/JKJP2
ADD . .
RUN go build ./cmd/mcproxy 

FROM debian:latest
RUN mkdir /home/vnf
COPY --from=builder /home/go/src/github.com/200106-uta-go/JKJP2 /home/vnf

RUN  mv /home/vnf/mcproxy /bin/
EXPOSE 80 443 22
EXPOSE 4444
WORKDIR /home/vnf