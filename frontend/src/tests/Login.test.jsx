import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import Login from '../pages/Login';
import { MemoryRouter } from 'react-router-dom';


vi.mock('../services/api', () => ({
login: vi.fn((email, pass) => {
if (email === 'test@test.com' && pass === '1234') return Promise.resolve({ token: '123' });
else return Promise.reject(new Error('Error'));
}),
}));


test('login exitoso guarda token y redirige', async () => {
render(<MemoryRouter><Login /></MemoryRouter>);
fireEvent.change(screen.getByPlaceholderText('Email'), { target: { value: 'test@test.com' } });
fireEvent.change(screen.getByPlaceholderText('Contraseña'), { target: { value: '1234' } });
fireEvent.click(screen.getByText('Entrar'));


await waitFor(() => expect(localStorage.getItem('token')).toBe('123'));
});


test('muestra error si las credenciales son inválidas', async () => {
render(<MemoryRouter><Login /></MemoryRouter>);
fireEvent.change(screen.getByPlaceholderText('Email'), { target: { value: 'bad@test.com' } });
fireEvent.change(screen.getByPlaceholderText('Contraseña'), { target: { value: '0000' } });
fireEvent.click(screen.getByText('Entrar'));


await waitFor(() => screen.getByText(/Credenciales inválidas/));
});