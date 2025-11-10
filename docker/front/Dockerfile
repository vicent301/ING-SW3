# ---- Build (solo assets estáticos) ----
FROM node:20-alpine AS builder
WORKDIR /app

# El contexto es frontend/, por eso estos archivos existen:
COPY package*.json ./
RUN npm ci

# (opcional) build arg para fallback de Vite
ARG VITE_API_URL=""
ENV VITE_API_URL=${VITE_API_URL}

COPY . .
RUN npm run build

# ---- Runtime (nginx + env.js inyectado al iniciar) ----
FROM nginx:alpine

# Nginx para SPA (el archivo debe estar en frontend/nginx.conf)
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Archivos estáticos compilados
COPY --from=builder /app/dist /usr/share/nginx/html

# Script de arranque que genera env.js según $API_URL (debe estar en frontend/entrypoint.sh)
COPY entrypoint.sh /docker-entrypoint.d/99-generate-env.sh
RUN chmod +x /docker-entrypoint.d/99-generate-env.sh

ENV API_URL=""
EXPOSE 80


