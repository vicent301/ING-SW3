#!/bin/sh
set -e

# Asegurar carpeta de assets
mkdir -p /usr/share/nginx/html/assets

# Generar env.js con la variable de entorno API_URL
cat > /usr/share/nginx/html/assets/env.js <<EOF
window.env = {
  apiUrl: "${API_URL}"
};
EOF

# Continuar con el entrypoint por defecto de nginx
# (nginx:alpine usa "nginx -g 'daemon off;'" como CMD)
