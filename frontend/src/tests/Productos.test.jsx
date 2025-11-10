// src/tests/Productos.test.jsx
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import { MemoryRouter, Routes, Route } from 'react-router-dom';
import { vi } from 'vitest';
import Productos from '../pages/Productos';

// Mock react-hot-toast (SIN variables externas)
vi.mock('react-hot-toast', () => {
  const toast = Object.assign(() => {}, {
    success: vi.fn(),
    error: vi.fn(),
  });
  return {
    __esModule: true,
    default: toast,
    Toaster: () => null,
  };
});

// Mock services/api
vi.mock('../services/api', async () => {
  const actual = await vi.importActual('../services/api');
  return {
    ...actual,
    getProducts: vi.fn(async () => [
      { id: 1, name: 'Nike Air', description: 'zapa', price: 150, image_url: 'img1.jpg' },
    ]),
    addToCart: vi.fn(async () => ({ success: true })),
  };
});

import toast from 'react-hot-toast';
import { addToCart } from '../services/api';

test('renderiza productos y permite agregar al carrito', async () => {
  render(
    <MemoryRouter initialEntries={['/products']}>
      <Routes>
        <Route path="/products" element={<Productos />} />
      </Routes>
    </MemoryRouter>
  );

  // aparece el producto mockeado
  expect(await screen.findByText('Nike Air')).toBeInTheDocument();

  // click en Agregar
  fireEvent.click(screen.getByRole('button', { name: /agregar al carrito/i }));

  // ⚠️ ambos asserts deben esperar al tick asíncrono
  await waitFor(() => expect(addToCart).toHaveBeenCalledWith(1, 1));
  await waitFor(() =>
    expect(toast.success).toHaveBeenCalledWith(expect.stringMatching(/Producto agregado/i))
  );
});
