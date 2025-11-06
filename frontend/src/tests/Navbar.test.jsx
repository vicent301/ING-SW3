import { render, screen, waitFor } from '@testing-library/react';
import Navbar from '../components/Navbar';
import { MemoryRouter } from 'react-router-dom';


vi.mock('../services/api', () => ({
getProfile: vi.fn(() => Promise.resolve({ name: 'Vicente Monzo' })),
getCart: vi.fn(() => Promise.resolve({ items: [1,2,3] })),
}));


test('muestra nombre del usuario y cantidad del carrito', async () => {
render(<MemoryRouter><Navbar /></MemoryRouter>);
await waitFor(() => screen.getByText(/Vicente Monzo/));
expect(screen.getByText('Vicente Monzo')).toBeInTheDocument();
expect(screen.getByText('3')).toBeInTheDocument();
});


// --- src/tests/ProtectedRoute.test.jsx ---
import { render } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import ProtectedRoute from '../components/ProtectedRoute';


test('redirige al login si no hay token', () => {
localStorage.removeItem('token');
const { container } = render(
<MemoryRouter>
<ProtectedRoute><div>Contenido protegido</div></ProtectedRoute>
</MemoryRouter>
);
expect(container.innerHTML).toContain('login');
});