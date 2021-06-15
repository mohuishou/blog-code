FROM golang:1.15 as builder

WORKDIR /app

COPY main.go main.go
COPY run.sh run.sh

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build main.go

FROM alpine:3

WORKDIR /app

COPY --from=builder /app /app

#COPY main.go main.go
COPY run.sh run.sh
#COPY nsswitch.conf /etc/nsswitch.conf

RUN chmod +x run.sh

ENV GODEBUG="netdns=go+2"
ENV CGO_ENABLED=0
ENV GOOS=linux

CMD /app/run.sh