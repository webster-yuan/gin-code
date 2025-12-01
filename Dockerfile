FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o gin-app main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/gin-app .
COPY config.yaml.example ./config.yaml

EXPOSE 8080

CMD ["./gin-app", "server"]