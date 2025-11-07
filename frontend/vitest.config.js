import { defineConfig } from 'vitest/config';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  test: {
    globals: true,
    environment: 'jsdom',
    setupFiles: './src/tests/setup.js',
    coverage: {
      provider: 'istanbul', // ✅ agrega el provider correcto
      reporter: ['text', 'lcov', 'cobertura', 'html'], // ✅ agrega los tipos de reporte
      reportsDirectory: 'coverage', // ✅ carpeta de salida
    },
    reporters: [
      'default',
      ['junit', { outputFile: 'test-results.xml' }], // para Azure
    ],
  },
});
