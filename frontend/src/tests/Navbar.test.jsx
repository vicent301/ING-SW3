import { render, screen, waitFor } from '@testing-library/react';
import Navbar from '../components/Navbar';
import { MemoryRouter } from 'react-router-dom';
import { vi } from 'vitest';

vi.mock('../services/api', () => ({
  getProfile: vi.fn(() => Promise.resolve({ name: 'Vicente Monzo' })),
  getCart: vi.fn(() => Promise.resolve({ items: [1, 2, 3] })),
}));

describe('Navbar', () => {
  test('muestra nombre del usuario y cantidad del carrito', async () => {
    render(
      <MemoryRouter>
        <Navbar />
      </MemoryRouter>
    );
    await waitFor(() => screen.getByText(/Vicente Monzo/i));
    expect(screen.getByText(/Vicente Monzo/i)).toBeInTheDocument();
    expect(screen.getByText('3')).toBeInTheDocument();
  });
});
