import { render, screen, fireEvent } from '@testing-library/react';
import Register from '../pages/Register';
import { MemoryRouter } from 'react-router-dom';


vi.mock('../services/api', () => ({
register: vi.fn(() => Promise.resolve({ success: true })),
}));


test('envía el formulario correctamente', async () => {
render(<MemoryRouter><Register /></MemoryRouter>);
fireEvent.change(screen.getByPlaceholderText('Nombre'), { target: { value: 'Juan' } });
fireEvent.change(screen.getByPlaceholderText('Email'), { target: { value: 'juan@test.com' } });
fireEvent.change(screen.getByPlaceholderText('Contraseña'), { target: { value: '1234' } });
fireEvent.click(screen.getByText('Crear Cuenta'));
});