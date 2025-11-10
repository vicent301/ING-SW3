// URL base del backend (sin /api)
const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080";


// 🛍️ Productos
export async function getProducts() {
  const res = await fetch(`${API_URL}/api/products`);
  if (!res.ok) throw new Error("Error al obtener productos");
  return res.json();
}

export async function getProductById(id) {
  const res = await fetch(`${API_URL}/api/products/${id}`);
  if (!res.ok) throw new Error("Error al obtener producto");
  return res.json();
}


// 🔐 Autenticación
export async function login(email, password) {
  const res = await fetch(`${API_URL}/api/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
  });
  if (!res.ok) throw new Error("Error al iniciar sesión");
  return res.json();
}

export async function register(userData) {
  const res = await fetch(`${API_URL}/api/register`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(userData),
  });
  if (!res.ok) throw new Error("Error al registrarse");
  return res.json();
}


// 👤 Perfil del usuario
export async function getProfile() {
  const token = localStorage.getItem("token");
  if (!token) throw new Error("No token");

  const res = await fetch(`${API_URL}/api/me`, {
    headers: { Authorization: `Bearer ${token}` },
  });

  if (!res.ok) throw new Error("Error al obtener perfil");
  return res.json();
}


// 🛒 Agregar producto al carrito
export async function addToCart(productId, quantity = 1) {
  const token = localStorage.getItem("token");
  if (!token) throw new Error("No token");

  const res = await fetch(`${API_URL}/api/cart/add`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify({ product_id: productId, quantity }),
  });

  if (!res.ok) throw new Error("Error al agregar producto al carrito");
  return res.json();
}


// 🛒 Obtener carrito
export async function getCart() {
  const token = localStorage.getItem("token");
  if (!token) throw new Error("No token");

  const res = await fetch(`${API_URL}/api/cart`, {
    headers: { Authorization: `Bearer ${token}` },
  });

  if (!res.ok) throw new Error("Error al obtener carrito");
  return res.json();
}


// ❌ Eliminar producto del carrito
export async function removeFromCart(productId) {
  const token = localStorage.getItem("token");
  if (!token) throw new Error("No token");

  const res = await fetch(`${API_URL}/api/cart/remove`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify({ product_id: productId }),
  });

  if (!res.ok) throw new Error("Error al eliminar producto");
  return res.json();
}


// 🧹 Vaciar carrito
export async function clearCart() {
  const token = localStorage.getItem("token");
  if (!token) throw new Error("No token");

  const res = await fetch(`${API_URL}/api/cart/clear`, {
    method: "DELETE",
    headers: { Authorization: `Bearer ${token}` },
  });

  if (!res.ok) throw new Error("Error al vaciar carrito");
  return res.json();
}
