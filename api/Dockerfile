FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev librdkafka-dev pkgconfig

COPY go.mod go.sum ./
COPY . .

ENV CGO_ENABLED=1
ENV PKG_CONFIG_PATH=/usr/lib/pkgconfig

# Use a tag de build para forçar a versão dinâmica
RUN go build -tags dynamic -v -o api ./main.go

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache librdkafka

COPY --from=builder /app/api /app/api

EXPOSE 8080

CMD ["/app/api"]