FROM golang:1.14.3-alpine3.11

WORKDIR app/

COPY app .

WORKDIR main

RUN go build main.go

CMD ["./main"]