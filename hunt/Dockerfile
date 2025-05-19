FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev librdkafka-dev pkgconfig

COPY go.mod go.sum ./
COPY . .

ENV CGO_ENABLED=1
ENV PKG_CONFIG_PATH=/usr/lib/pkgconfig

# Use a tag de build para forçar a versão dinâmica
RUN go build -tags dynamic -v -o hunt ./main.go

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache librdkafka

COPY --from=builder /app/hunt /app/hunt

CMD ["/app/hunt"]