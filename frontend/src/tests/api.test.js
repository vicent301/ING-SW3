import { vi, describe, it, beforeAll, afterAll, beforeEach, expect } from 'vitest';
import {
  getProducts,
  getProductById,
  login,
  register,
  getProfile,
  addToCart,
  getCart,
  removeFromCart,
  clearCart,
} from '../services/api';

// Creamos un mock global de fetch y lo reutilizamos en todo el suite
let fetchMock;

beforeAll(() => {
  fetchMock = vi.fn();
  // reemplaza el fetch nativo por un mock de Vitest
  vi.stubGlobal('fetch', fetchMock);
});

afterAll(() => {
  // restaura todos los globals stubbeados
  vi.unstubAllGlobals();
});

beforeEach(() => {
  fetchMock.mockReset();
  localStorage.clear();
});

describe('services/api.js', () => {
  const TOKEN = 'tok_123';

  // --------------------
  // Productos
  // --------------------
  it('getProducts() retorna items (200 OK)', async () => {
    fetchMock.mockResolvedValueOnce({
      ok: true,
      json: async () => ({ items: [{ id: 1, name: 'Zapa' }] }),
    });

    const res = await getProducts();
    expect(res.items).toHaveLength(1);
    expect(fetchMock).toHaveBeenCalledTimes(1);
  });

  it('getProductById() retorna un producto', async () => {
    fetchMock.mockResolvedValueOnce({
      ok: true,
      json: async () => ({ id: 99, name: 'Prod-99' }),
    });

    const res = await getProductById(99);
    expect(res.id).toBe(99);
    expect(fetchMock).toHaveBeenCalledTimes(1);
  });

  it('getProducts() lanza error cuando !ok', async () => {
    fetchMock.mockResolvedValueOnce({ ok: false });
    await expect(getProducts()).rejects.toThrow(/obtener productos/i);
  });

  // --------------------
  // Auth
  // --------------------
  it('login() hace POST y devuelve token', async () => {
    fetchMock.mockResolvedValueOnce({
      ok: true,
      json: async () => ({ token: 'abc' }),
    });

    const res = await login('test@example.com', '123456');
    expect(res.token).toBe('abc');
    expect(fetchMock).toHaveBeenCalledWith(
      expect.stringMatching(/\/api\/login$/),
      expect.objectContaining({
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email: 'test@example.com', password: '123456' }),
      })
    );
  });

  it('register(userData) hace POST y devuelve user', async () => {
    const payload = { name: 'Vicente', email: 'v@ex.com', password: '123' };

    fetchMock.mockResolvedValueOnce({
      ok: true,
      json: async () => ({ id: 1, ...payload }),
    });

    const res = await register(payload);
    expect(res.email).toBe('v@ex.com');
    expect(fetchMock).toHaveBeenCalledWith(
      expect.stringMatching(/\/api\/register$/),
      expect.objectContaining({
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
      })
    );
  });

  // --------------------
  // Perfil (requiere token)
  // --------------------
  it('getProfile() devuelve usuario cuando hay token', async () => {
    localStorage.setItem('token', TOKEN);

    fetchMock.mockResolvedValueOnce({
      ok: true,
      json: async () => ({ name: 'Simón' }),
    });

    const res = await getProfile();
    expect(res.name).toBe('Simón');
    expect(fetchMock).toHaveBeenCalledWith(
      expect.stringMatching(/\/api\/me$/),
      expect.objectContaining({
        headers: expect.objectContaining({ Authorization: `Bearer ${TOKEN}` }),
      })
    );
  });

  it('getProfile() lanza "No token" si no hay token', async () => {
    await expect(getProfile()).rejects.toThrow(/no token/i);
  });

  // --------------------
  // Carrito (requiere token)
  // --------------------
  it('addToCart() agrega y devuelve success', async () => {
    localStorage.setItem('token', TOKEN);

    fetchMock.mockResolvedValueOnce({
      ok: true,
      json: async () => ({ success: true }),
    });

    const res = await addToCart(7, 2);
    expect(res.success).toBe(true);
    expect(fetchMock).toHaveBeenCalledWith(
      expect.stringMatching(/\/api\/cart\/add$/),
      expect.objectContaining({
        method: 'POST',
        headers: expect.objectContaining({
          Authorization: `Bearer ${TOKEN}`,
          'Content-Type': 'application/json',
        }),
        body: JSON.stringify({ product_id: 7, quantity: 2 }),
      })
    );
  });

  it('getCart() devuelve items', async () => {
    localStorage.setItem('token', TOKEN);

    fetchMock.mockResolvedValueOnce({
      ok: true,
      json: async () => ({ items: [{ id: 1, quantity: 1 }] }),
    });

    const res = await getCart();
    expect(res.items).toHaveLength(1);
  });

  it('removeFromCart() elimina ítem', async () => {
    localStorage.setItem('token', TOKEN);

    fetchMock.mockResolvedValueOnce({
      ok: true,
      json: async () => ({ removed: true }),
    });

    const res = await removeFromCart(1);
    expect(res.removed).toBe(true);
    expect(fetchMock).toHaveBeenCalledWith(
      expect.stringMatching(/\/api\/cart\/remove$/),
      expect.objectContaining({
        method: 'DELETE',
        headers: expect.objectContaining({
          Authorization: `Bearer ${TOKEN}`,
          'Content-Type': 'application/json',
        }),
        body: JSON.stringify({ product_id: 1 }),
      })
    );
  });

  it('clearCart() vacía el carrito', async () => {
    localStorage.setItem('token', TOKEN);

    fetchMock.mockResolvedValueOnce({
      ok: true,
      json: async () => ({ cleared: true }),
    });

    const res = await clearCart();
    expect(res.cleared).toBe(true);
    expect(fetchMock).toHaveBeenCalledWith(
      expect.stringMatching(/\/api\/cart\/clear$/),
      expect.objectContaining({
        method: 'DELETE',
        headers: expect.objectContaining({
          Authorization: `Bearer ${TOKEN}`,
        }),
      })
    );
  });

  // --------------------
  // Ramas negativas
  // --------------------
  it('addToCart() lanza error cuando !ok', async () => {
    localStorage.setItem('token', TOKEN);
    fetchMock.mockResolvedValueOnce({ ok: false });
    await expect(addToCart(1)).rejects.toThrow(/agregar producto/i);
  });

  it('getCart() lanza error cuando !ok', async () => {
    localStorage.setItem('token', TOKEN);
    fetchMock.mockResolvedValueOnce({ ok: false });
    await expect(getCart()).rejects.toThrow(/obtener carrito/i);
  });
});
