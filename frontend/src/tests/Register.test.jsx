import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import Register from '../pages/Register';
import { vi } from 'vitest';
import { MemoryRouter } from 'react-router-dom';

const mockNavigate = vi.fn();
vi.mock('react-router-dom', async () => {
  const actual = await vi.importActual('react-router-dom');
  return {
    ...actual,
    useNavigate: () => mockNavigate,
  };
});

vi.mock('../services/api', () => ({
  register: vi.fn(() => Promise.resolve({ success: true })),
}));

describe('Register', () => {
  test('envía el formulario y navega a /login', async () => {
    render(
      <MemoryRouter>
        <Register />
      </MemoryRouter>
    );

    fireEvent.change(screen.getByPlaceholderText('Nombre'), { target: { value: 'Juan' } });
    fireEvent.change(screen.getByPlaceholderText('Email'), { target: { value: 'juan@test.com' } });
    fireEvent.change(screen.getByPlaceholderText('Contraseña'), { target: { value: '1234' } });
    fireEvent.click(screen.getByText('Crear Cuenta'));

    await waitFor(() => expect(mockNavigate).toHaveBeenCalledWith('/login'));
  });
});
