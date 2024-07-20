FROM golang:alpine AS builder
WORKDIR /go/src
COPY . .
RUN go build -o api server.go

FROM alpine
COPY  --from=builder go/src/app .

ENTRYPOINT ["./api"]