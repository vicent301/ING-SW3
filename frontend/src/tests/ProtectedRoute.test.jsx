import { render, screen } from '@testing-library/react';
import { MemoryRouter, Routes, Route } from 'react-router-dom';
import ProtectedRoute from '../components/ProtectedRoute';

test('redirige al login si no hay token', () => {
  localStorage.removeItem('token');

  render(
    <MemoryRouter initialEntries={['/privado']}>
      <Routes>
        <Route path="/login" element={<div>Pantalla Login</div>} />
        <Route
          path="/privado"
          element={
            <ProtectedRoute>
              <div>Privado</div>
            </ProtectedRoute>
          }
        />
      </Routes>
    </MemoryRouter>
  );

  expect(screen.getByText(/Pantalla Login/i)).toBeInTheDocument();
});
