FROM golang:1.25-bookworm AS builder

WORKDIR /app

COPY . .

WORKDIR /app/services/Log

RUN go mod download

RUN CGO_ENABLED=0 go build -o ./cmd/exe ./cmd

FROM alpine:3.23

WORKDIR /app

COPY --from=builder /app/services/Log/cmd/exe /app

CMD ["/app/exe"]