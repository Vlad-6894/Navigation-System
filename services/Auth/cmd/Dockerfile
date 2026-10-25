FROM golang:1.26-bookworm AS builder

WORKDIR /app

COPY go.work ./

COPY pkg/go.mod pkg/go.su* ./pkg/
COPY navigation_proto/go.mod navigation-proto/go.su* ./navigation-proto/
COPY services/Auth/go.mod services/Auth/go.su* ./services/Auth/

RUN go mod download

COPY pkg/ ./pkg/
COPY navigation_proto/ ./navigation_proto/
COPY services/Auth/ ./services/Auth/

RUN CGO_ENABLED=0 go build -o /app/services/Auth/cmd/exe ./services/Auth/cmd

FROM alpine:3.23

WORKDIR /app

COPY --from=builder /app/services/Auth/cmd/exe /app

CMD ["/app/exe"]