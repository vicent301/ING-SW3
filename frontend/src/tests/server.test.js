import { describe, it, expect, beforeAll, afterAll, afterEach } from 'vitest';
import { server } from '../tests/server';
import { http, HttpResponse } from 'msw';

const API_URL = import.meta.env.VITE_API_URL || '/api';

beforeAll(() => server.listen());
afterEach(() => server.resetHandlers());
afterAll(() => server.close());

describe("MSW Server - Handlers", () => {

  // ✅ GET /products
  it("retorna lista de productos", async () => {
    const res = await fetch(`${API_URL}/api/products`);
    const data = await res.json();

    expect(res.ok).toBe(true);
    expect(Array.isArray(data)).toBe(true);
    expect(data[0]).toHaveProperty('name');
  });

  // ✅ GET /cart
  it("retorna carrito con un item mockeado", async () => {
    const res = await fetch(`${API_URL}/api/cart`);
    const data = await res.json();

    expect(res.ok).toBe(true);
    expect(data.items.length).toBe(1);
    expect(data.items[0]).toHaveProperty('product_id', 1);
  });

  // ✅ POST /login (success)
  it("login correcto devuelve un token", async () => {
    const res = await fetch(`${API_URL}/api/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email: "test@test.com", password: "1234" }),
    });

    const data = await res.json();

    expect(res.ok).toBe(true);
    expect(data).toHaveProperty("token", "fake-jwt-token");
  });

  // ✅ POST /login (error)
  it("login incorrecto debería retornar 401", async () => {
    const res = await fetch(`${API_URL}/api/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email: "wrong@mail.com", password: "bad" }),
    });

    expect(res.status).toBe(401);
  });

  // ✅ GET /me
  it("retorna perfil del usuario", async () => {
    const res = await fetch(`${API_URL}/api/me`);
    const data = await res.json();

    expect(res.ok).toBe(true);
    expect(data).toHaveProperty("name", "Vicente Monzo");
  });

  // ✅ Simulación de error en /products para cubrir 100% ramas
  it("maneja error del servidor en /products", async () => {
    server.use(
      http.get(`${API_URL}/api/products`, () =>
        HttpResponse.text("Server error", { status: 500 })
      )
    );

    const res = await fetch(`${API_URL}/api/products`);
    expect(res.ok).toBe(false);
    expect(res.status).toBe(500);
  });
});
