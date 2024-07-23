FROM alpine AS builder
RUN apk update && apk add go 
COPY . .
RUN go build -o consumer main.go

FROM alpine
COPY  --from=builder /consumer .

ENTRYPOINT ["./consumer"]