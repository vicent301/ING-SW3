// Lee la URL desde la variable de entorno en tiempo de build
const API_URL = "https://appweb-api-prod-hdhgb2bmb6eyaubv.chilecentral-01.azurewebsites.net/api";

// 🛍️ Productos
export async function getProducts() {
  const res = await fetch(`${API_URL}/products`);
  if (!res.ok) throw new Error("Error al obtener productos");
  return res.json();
}

export async function getProductById(id) {
  const res = await fetch(`${API_URL}/products/${id}`);
  if (!res.ok) throw new Error("Error al obtener producto");
  return res.json();
}

// 🔐 Autenticación
export async function login(email, password) {
  const res = await fetch(`${API_URL}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
  });
  if (!res.ok) throw new Error("Error al iniciar sesión");
  return res.json();
}

export async function register(userData) {
  const res = await fetch(`${API_URL}/register`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(userData),
  });
  if (!res.ok) throw new Error("Error al registrarse");
  return res.json();
}

// 👤 Perfil del usuario logueado
export async function getProfile() {
  const token = localStorage.getItem("token");
  if (!token) throw new Error("No token");
  
  const res = await fetch(`${API_URL}/me`, {
    headers: { Authorization: `Bearer ${token}` },
  });
  if (!res.ok) throw new Error("Error al obtener perfil");
  return res.json();
}



//agrega producto al carrito
export async function addToCart(productId, quantity = 1) {
  const token = localStorage.getItem("token");
  if (!token) throw new Error("No token");

  const res = await fetch(`${API_URL}/cart/add`, {
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
// 🛒 Obtener el carrito actual
export async function getCart() {
  const token = localStorage.getItem("token");
  if (!token) throw new Error("No token");

  const res = await fetch(`${API_URL}/cart`, {
    headers: { Authorization: `Bearer ${token}` },
  });
  if (!res.ok) throw new Error("Error al obtener carrito");
  return res.json();
}

// ❌ Eliminar un producto del carrito
export async function removeFromCart(productId) {
  const token = localStorage.getItem("token");
  if (!token) throw new Error("No token");

  const res = await fetch(`${API_URL}/cart/remove`, {
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

// 🧹 Vaciar el carrito completo
export async function clearCart() {
  const token = localStorage.getItem("token");
  if (!token) throw new Error("No token");

  const res = await fetch(`${API_URL}/cart/clear`, {
    method: "DELETE",
    headers: { Authorization: `Bearer ${token}` },
  });

  if (!res.ok) throw new Error("Error al vaciar carrito");
  return res.json();
}
