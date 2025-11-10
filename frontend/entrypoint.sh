#!/bin/sh
set -e

# Generar /env.js con la variable de entorno API_URL
cat > /usr/share/nginx/html/env.js <<EOF
window.__ENV = {
  API_BASE: "${API_URL}"
};
EOF

# Arrancar nginx en primer plano
exec nginx -g 'daemon off;'
