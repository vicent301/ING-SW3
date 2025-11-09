# ---- Build (solo assets estáticos) ----
FROM node:20-alpine AS builder
WORKDIR /app

COPY package*.json ./
RUN npm ci

# OJO: las VITE_ quedan para fallback local; en runtime usaremos env.js
ARG VITE_API_URL=""
ENV VITE_API_URL=${VITE_API_URL}

COPY . .
RUN npm run build

# ---- Runtime (nginx + env.js inyectado al iniciar) ----
FROM nginx:alpine

# Nginx para SPA
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Archivos estáticos compilados
COPY --from=builder /app/dist /usr/share/nginx/html

# Script de arranque que genera env.js según $API_URL
COPY entrypoint.sh /docker-entrypoint.d/99-generate-env.sh
RUN chmod +x /docker-entrypoint.d/99-generate-env.sh

ENV API_URL=""
EXPOSE 80

# Healthcheck simple
HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
  CMD wget -qO- http://127.0.0.1/ > /dev/null || exit 1
