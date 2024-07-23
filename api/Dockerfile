FROM alpine AS builder
RUN apk update && apk add go 
COPY . .
RUN go build -o app server.go

FROM alpine
COPY  --from=builder /app .

ENTRYPOINT ["./app"]