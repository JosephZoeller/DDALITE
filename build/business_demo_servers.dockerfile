FROM golang:rc-buster AS builder
RUN mkdir -p /home/go/src/github.com/200106-uta-go/JKJP2
ADD . /home/go/src/github.com/200106-uta-go/JKJP2
WORKDIR /home/go/src/github.com/200106-uta-go/JKJP2

RUN go build -o business_demo_servers .

FROM debian:latest

COPY --from=builder /home/go/src/github.com/200106-uta-go/JKJP2 /home/servers/
WORKDIR /home/servers/
EXPOSE 80 443 22
EXPOSE 8081 8082 8083
EXPOSE 9090
CMD [ "./business_demo_servers" ]
