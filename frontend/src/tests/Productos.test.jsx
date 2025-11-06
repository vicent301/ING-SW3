import { render, screen, waitFor, fireEvent } from '@testing-library/react';
import Productos from '../pages/Productos';
import { vi } from 'vitest';


vi.mock('../services/api', async () => {
const actual = await vi.importActual('../services/api');
return {
...actual,
getProducts: vi.fn(() => Promise.resolve([
{ id: 1, name: 'Nike Air', price: 150, image_url: 'img1.jpg' },
])),
addToCart: vi.fn(() => Promise.resolve({ success: true })),
};
});


test('renderiza productos y permite agregar al carrito', async () => {
render(<Productos />);
expect(screen.getByText(/Cargando productos/)).toBeInTheDocument();


await waitFor(() => screen.getByText('Nike Air'));
expect(screen.getByText('Nike Air')).toBeInTheDocument();


const button = screen.getByText(/Agregar al carrito/);
fireEvent.click(button);
await waitFor(() => screen.getByText(/Producto agregado al carrito/));
});