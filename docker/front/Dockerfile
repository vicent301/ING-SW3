==== Build stage ====
FROM node:20-alpine AS builder
WORKDIR /app

COPY package*.json ./
RUN npm ci

COPY . .

IMPORTANTE:
Tu api.js hace: const API_URL = import.meta.env.VITE_API_URL || "/api";
y luego usa ${API_URL}/api/...
=> VITE_API_URL debe NO incluir /api (ej: https://appweb-api-qa-xxx.azurewebsites.net/)
ARG VITE_API_URL
ENV VITE_API_URL=${VITE_API_URL}

RUN npm run build

==== Runtime stage (Nginx) ====
FROM nginx:alpine
Reemplazamos el default con nuestra config SPA
COPY nginx.conf /etc/nginx/conf.d/default.conf

Copiamos artefactos
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

Opcional: healthcheck HTTP simple a /
HEALTHCHECK --interval=30s --timeout=5s --retries=3 CMD wget -qO- http://127.0.0.1/ || exit 1
frontend