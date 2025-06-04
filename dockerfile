# Etapa 1: Build
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o countries-backend main.go

# Etapa 2: Producción
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/countries-backend .

EXPOSE 8080

CMD ["./countries-backend"]
