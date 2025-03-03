FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download


COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/microservice ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/microservice /app/microservice

EXPOSE 8000

CMD ["./microservice"]