# ===== ETAPA DE BUILD =====
FROM node:20-alpine AS build
WORKDIR /app

COPY package*.json ./
RUN npm ci

# 👇 ENCHUFAMOS LA VARIABLE PARA VITE
ARG VITE_API_BASE=""
ENV VITE_API_BASE=${VITE_API_BASE}

COPY . .

# Vite va a leer import.meta.env.VITE_API_BASE aquí
RUN npm run build

# ===== ETAPA DE RUN =====
FROM node:20-alpine
WORKDIR /app

RUN npm install -g serve

COPY --from=build /app/dist ./dist

ENV PORT=4173
EXPOSE 4173

CMD ["sh", "-c", "serve -s dist -l ${PORT}"]
