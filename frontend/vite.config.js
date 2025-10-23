import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      "/api": "https://appweb-api-prod-hdhgb2bmb6eyaubv.chilecentral-01.azurewebsites.net",
    },
    historyApiFallback: true, // 👈 esto asegura que las rutas vuelvan a index.html
  },
});
