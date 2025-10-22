import { useEffect, useState } from "react";
import { getProducts, addToCart } from "../services/api";

export default function Productos() {
  const [productos, setProductos] = useState([]);
  const [loading, setLoading] = useState(true);
  const [msg, setMsg] = useState("");

  useEffect(() => {
    async function fetchData() {
      try {
        const data = await getProducts();
        setProductos(data);
      } catch (err) {
        console.error(err);
      } finally {
        setLoading(false);
      }
    }
    fetchData();
  }, []);

  const handleAddToCart = async (id) => {
    try {
      await addToCart(id);
      setMsg("✅ Producto agregado al carrito");
      setTimeout(() => setMsg(""), 2000);
    } catch (err) {
      setMsg("⚠️ Debes iniciar sesión para agregar al carrito");
      setTimeout(() => setMsg(""), 2000);
    }
  };

  if (loading) return <p className="text-center mt-10">Cargando productos...</p>;

  return (
    <div className="p-10 grid grid-cols-1 md:grid-cols-3 gap-6">
      {productos.map((p) => (
        <div
          key={p.id}
          className="border rounded-xl shadow hover:shadow-lg p-4 transition"
        >
          <img
            src={p.image_url || "https://via.placeholder.com/200"}
            alt={p.name}
            className="w-full h-48 object-cover rounded-md"
          />
          <h2 className="text-xl font-semibold mt-3">{p.name}</h2>
          <p className="text-gray-600 mb-3">${p.price}</p>

          <button
            onClick={() => handleAddToCart(p.id)}
            className="w-full bg-black text-white py-2 rounded hover:bg-gray-800 transition"
          >
            Agregar al carrito 🛒
          </button>
        </div>
      ))}

      {msg && (
        <div className="fixed bottom-4 right-4 bg-black text-white px-4 py-2 rounded shadow-lg text-sm">
          {msg}
        </div>
      )}
    </div>
  );
}
