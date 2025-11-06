import { setupServer } from 'msw/node';
import { http, HttpResponse } from 'msw';


const API_URL = import.meta.env.VITE_API_URL || '/api';


export const handlers = [
http.get(`${API_URL}/api/products`, () =>
HttpResponse.json([
{ id: 1, name: 'Nike Air', price: 150, image_url: 'img1.jpg' },
{ id: 2, name: 'Adidas Samba', price: 120, image_url: 'img2.jpg' },
])
),
http.get(`${API_URL}/api/cart`, () =>
HttpResponse.json({ items: [{ id: 1, product_id: 1, quantity: 2, product: { name: 'Nike Air', price: 150 } }] })
),
http.post(`${API_URL}/api/login`, async ({ request }) => {
const body = await request.json();
if (body.email === 'test@test.com' && body.password === '1234') {
return HttpResponse.json({ token: 'fake-jwt-token' });
}
return new HttpResponse('Unauthorized', { status: 401 });
}),
http.get(`${API_URL}/api/me`, () =>
HttpResponse.json({ id: 1, name: 'Vicente Monzo', email: 'vicente@ucc.com' })
),
];


export const server = setupServer(...handlers);