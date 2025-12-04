# ===== ETAPA DE BUILD =====
FROM node:20-alpine AS build
WORKDIR /app

# Dependencias
COPY package*.json ./
RUN npm ci

# Si usás VITE_API_URL en el código, la pasamos como ARG
ARG VITE_API_URL=""
ENV VITE_API_URL=${VITE_API_URL}

# Código fuente
COPY . .

# Build de producción (Vite)
RUN npm run build

# ===== ETAPA DE RUN =====
FROM node:20-alpine
WORKDIR /app

# Servidor estático para los archivos de Vite
RUN npm install -g serve

# Copiamos el build
COPY --from=build /app/dist ./dist

# Railway setea PORT, pero por si acaso dejamos un default
ENV PORT=4173

EXPOSE 4173

# IMPORTANTÍSIMO: escuchar en ${PORT}, que es lo que usa Railway
CMD ["sh", "-c", "serve -s dist -l ${PORT}"]
