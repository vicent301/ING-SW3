import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { getCart, removeFromCart, clearCart } from "../services/api";

export default function Carrito() {
  const [cart, setCart] = useState([]);
  const [loading, setLoading] = useState(true);
  const [msg, setMsg] = useState("");

  useEffect(() => {
    async function loadCart() {
      try {
        const data = await getCart();
        console.log("🛒 Respuesta del backend:", data);
        setCart(data.items || []); // porque el backend devuelve { items: [...] }
      } catch (err) {
        setMsg("⚠️ Debes iniciar sesión para ver el carrito");
      } finally {
        setLoading(false);
      }
    }
    loadCart();
  }, []);

  const handleRemove = async (id) => {
    try {
      await removeFromCart(id);
      setCart(cart.filter((item) => item.product_id !== id));
    } catch {
      setMsg("Error al eliminar producto");
    }
  };

  const handleClear = async () => {
    try {
      await clearCart();
      setCart([]);
    } catch {
      setMsg("Error al vaciar carrito");
    }
  };

  if (loading) return <p className="text-center mt-10">Cargando carrito...</p>;

  if (!cart.length)
    return (
      <div className="text-center py-20">
        <h2 className="text-2xl font-bold mb-4">🛍️ Tu carrito está vacío</h2>
        <Link
          to="/products"
          className="text-blue-600 hover:underline font-semibold"
        >
          Ver productos
        </Link>
      </div>
    );

  const total = cart.reduce(
    (acc, item) => acc + (item.product?.price || 0) * (item.quantity || 1),
    0
  );

  return (
    <div className="max-w-4xl mx-auto bg-white shadow-md rounded-lg p-8 mt-10">
      <h1 className="text-3xl font-bold mb-6 text-center">🛒 Tu Carrito</h1>

      <div className="divide-y divide-gray-300">
        {cart.map((item) => (
          <div
            key={item.id}
            className="flex items-center justify-between py-4"
          >
            <div className="flex items-center gap-4">
              <img
                src={item.product?.image_url || "https://via.placeholder.com/80"}
                alt={item.product?.name}
                className="w-20 h-20 object-cover rounded-md"
              />
              <div>
                <h3 className="text-lg font-semibold">
                  {item.product?.name}
                </h3>
                <p className="text-sm text-gray-600">
                  ${item.product?.price} x {item.quantity || 1}
                </p>
              </div>
            </div>

            <button
              onClick={() => handleRemove(item.product_id)}
              className="text-red-600 hover:text-red-800 font-semibold"
            >
              Quitar
            </button>
          </div>
        ))}
      </div>

      <div className="mt-6 flex justify-between items-center">
        <button
          onClick={handleClear}
          className="bg-gray-200 hover:bg-gray-300 text-gray-800 px-4 py-2 rounded"
        >
          Vaciar carrito
        </button>

        <h2 className="text-xl font-bold">Total: ${total.toFixed(2)}</h2>
      </div>

      {msg && (
        <div className="mt-4 text-center text-sm text-red-600">{msg}</div>
      )}
    </div>
  );
}
