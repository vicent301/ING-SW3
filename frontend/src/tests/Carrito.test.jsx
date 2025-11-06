import { render, screen, waitFor, fireEvent } from '@testing-library/react';
import Carrito from '../pages/Carrito';
import { vi } from 'vitest';


vi.mock('../services/api', () => ({
getCart: vi.fn(() => Promise.resolve({ items: [ { id: 1, product_id: 1, quantity: 2, product: { name: 'Nike Air', price: 150 } } ] })),
removeFromCart: vi.fn(() => Promise.resolve({ success: true })),
clearCart: vi.fn(() => Promise.resolve({ success: true })),
}));


test('renderiza carrito y permite vaciarlo', async () => {
render(<Carrito />);
await waitFor(() => screen.getByText(/Tu Carrito/));
expect(screen.getByText('Nike Air')).toBeInTheDocument();


fireEvent.click(screen.getByText(/Vaciar carrito/));
await waitFor(() => screen.getByText(/Tu carrito está vacío/));
});


// --- src/tests/Navbar.test.jsx ---
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