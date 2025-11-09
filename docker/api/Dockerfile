FROM golang:1.21-alpine AS builder
WORKDIR /app

Dependencias
COPY go.mod go.sum ./
RUN go mod download

Código
COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o server ./...

FROM alpine:3.20
WORKDIR /app

Usuario no-root
RUN adduser -D -u 10001 appuser
USER appuser

Copiamos binario
COPY --from=builder /app/server /app/server

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --retries=3 CMD wget -qO- http://127.0.0.1:8080/api/healthz || exit 1

CMD ["/app/server"]