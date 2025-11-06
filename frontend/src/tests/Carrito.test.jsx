import { render, screen, waitFor, fireEvent } from '@testing-library/react';
import Carrito from '../pages/Carrito';
import { vi } from 'vitest';
import { MemoryRouter } from 'react-router-dom';

const mockGetCart = vi.fn(() =>
  Promise.resolve({
    items: [{ id: 1, product_id: 1, quantity: 2, product: { name: 'Nike Air', price: 150 } }],
  })
);
const mockRemove = vi.fn(() => Promise.resolve({ success: true }));
const mockClear = vi.fn(() => Promise.resolve({ success: true }));

vi.mock('../services/api', () => ({
  getCart: (...args) => mockGetCart(...args),
  removeFromCart: (...args) => mockRemove(...args),
  clearCart: (...args) => mockClear(...args),
}));

describe('Carrito', () => {
  test('renderiza carrito y calcula total', async () => {
    render(
      <MemoryRouter>
        <Carrito />
      </MemoryRouter>
    );

    expect(await screen.findByText(/Tu Carrito/i)).toBeInTheDocument();
    expect(screen.getByText('Nike Air')).toBeInTheDocument();
    // 150 * 2 = 300
    expect(screen.getByText(/Total: \$300\.00/i)).toBeInTheDocument();
  });

  test('permite vaciar el carrito', async () => {
    render(
      <MemoryRouter>
        <Carrito />
      </MemoryRouter>
    );

    await screen.findByText(/Tu Carrito/i);
    fireEvent.click(screen.getByText(/Vaciar carrito/i));

    // Esperá a que el estado se actualice y renderice el empty state
    await waitFor(() =>
      expect(screen.getByText(/Tu carrito está vacío/i)).toBeInTheDocument()
    );
  });
});
