# ---- Build stage ----
FROM golang:1.21-alpine AS builder
WORKDIR /app

# Dependencias
COPY go.mod go.sum ./
RUN go mod download

# Código
COPY . .

# Compilar binario estático para Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o appweb-backend ./...

# ---- Runtime stage ----
FROM alpine:3.20
# Certificados para TLS (MySQL, etc.)
RUN apk add --no-cache ca-certificates
WORKDIR /app

# Copiar binario
COPY --from=builder /app/appweb-backend /app/appweb-backend

# Puerto de tu API (tu app usa 8080)
ENV PORT=8080
EXPOSE 8080

# Usuario no root (opcional)
# RUN adduser -D -g '' appuser && chown -R appuser:appuser /app
# USER appuser

# Ejecutar
CMD ["/app/appweb-backend"]
