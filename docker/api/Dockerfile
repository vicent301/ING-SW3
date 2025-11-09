# docker/api/Dockerfile
# ===== Build stage =====
FROM golang:1.24.8-alpine AS builder
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOTOOLCHAIN=auto
WORKDIR /app

# deps necesarias para go mod y certificados
RUN apk add --no-cache git ca-certificates

# Copiamos módulos primero para cachear mejor
COPY go.mod go.sum ./
RUN go mod download

# Ahora el resto del código (el contexto del build es la carpeta backend)
COPY . .

# Compilamos binario estático
RUN go build -ldflags "-s -w" -o appweb-backend ./cmd/... || go build -ldflags "-s -w" -o appweb-backend .

# ===== Runtime stage =====
FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/appweb-backend /app/appweb-backend

EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=5s --retries=3 CMD wget -qO- http://127.0.0.1:8080/healthz >/dev/null 2>&1 || exit 1

ENTRYPOINT ["/app/appweb-backend"]
