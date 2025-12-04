# ==== build stage ====
FROM node:20-alpine AS build
WORKDIR /app

# copiar package.json y lock
COPY package*.json ./

RUN npm ci

# copiar el resto del código
COPY . .

# build de producción
RUN npm run build

# ==== production stage ====
FROM nginx:alpine
WORKDIR /usr/share/nginx/html

# borrar contenido default de nginx
RUN rm -rf ./*

# copiar build generado
COPY --from=build /app/dist .

# copiar configuración custom si la tenés (opcional)
# COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
