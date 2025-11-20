FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY vendor ./vendor

COPY . .

RUN go build -mod=vendor -o main .


FROM alpine:latest

WORKDIR /root/

RUN apk add --no-cache tzdata

COPY --from=builder /app/main .
COPY --from=builder /app/.env . 

EXPOSE 3000

CMD ["./main"]