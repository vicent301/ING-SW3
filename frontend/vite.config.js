import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      "/api": "http://localhost:8080",
    },
    historyApiFallback: true, // 👈 esto asegura que las rutas vuelvan a index.html
  },
});
