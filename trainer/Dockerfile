FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o trainer ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/trainer /app/trainer
EXPOSE 50051

CMD ["/app/trainer"]