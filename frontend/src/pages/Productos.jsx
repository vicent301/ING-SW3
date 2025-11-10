import { useEffect, useMemo, useState } from "react";
import { useLocation } from "react-router-dom";
import logo from "../assets/logo.png";
import toast from "react-hot-toast";

export default function Productos() {
  const [productos, setProductos] = useState([]);
  const location = useLocation();

  // ✅ leer "search" desde la URL
  const params = new URLSearchParams(location.search);
  const search = params.get("search")?.toLowerCase() || "";

  // ✅ formateador de precio
  const priceFmt = useMemo(
    () =>
      new Intl.NumberFormat("es-AR", {
        style: "currency",
        currency: "ARS",
        minimumFractionDigits: 0,
      }),
    []
  );

  // ✅ cargar productos
  useEffect(() => {
    async function cargarProductos() {
      try {
        const resp = await fetch("http://localhost:8080/api/products");
        const data = await resp.json();
        setProductos(Array.isArray(data) ? data : []);
      } catch (error) {
        console.error("Error cargando productos:", error);
      }
    }
    cargarProductos();
  }, []);

  // ✅ aplicar filtro del buscador
  const productosFiltrados = productos.filter(
    (p) =>
      p.name?.toLowerCase().includes(search) ||
      p.description?.toLowerCase().includes(search)
  );

  // ✅ agregar al carrito
  async function agregarAlCarrito(productId) {
    const token = localStorage.getItem("token");

    if (!token) {
      toast.success("✅ Producto agregado al carrito");
      return;
    }

    try {
      const resp = await fetch("http://localhost:8080/api/cart/add", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          product_id: productId,
          quantity: 1,
        }),
      });

      const data = await resp.json();

      if (!resp.ok) {
        alert("Error agregando al carrito: " + (data.error || "desconocido"));
        return;
      }

      toast.success("✅ Producto agregado al carrito");
    } catch (error) {
      console.error("Error al agregar al carrito:", error);
    }
  }

  return (
    <div className="w-full">
      {/* ✅ Logo (más cerca del navbar y tamaño válido) */}
      <div className="flex justify-center mt-6 mb-2">
        <img src={logo} alt="ZapaStore logo" className="w-20 opacity-90" />
      </div>

      {/* ✅ Título (menos separación) */}
      <h2 className="text-center text-3xl font-bold mt-1 mb-8 tracking-wide">
        Catálogo de Productos
      </h2>

      {/* ✅ Sin resultados */}
      {search && productosFiltrados.length === 0 && (
        <p className="text-center text-gray-500 mb-8">
          No se encontraron resultados para <span className="font-semibold">“{params.get("search")}”</span>.
        </p>
      )}

      {/* ✅ Grid */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-10 max-w-7xl mx-auto px-6 pb-12">
        {productosFiltrados.map((p) => (
          <article
            key={p.id}
            className="bg-white border border-gray-200 rounded-xl shadow-md hover:shadow-xl transition p-4"
          >
            <div className="w-full h-56 bg-gray-100 rounded-lg overflow-hidden">
              <img
                src={p.image_url || "https://via.placeholder.com/600x400?text=Producto"}
                alt={p.name}
                className="w-full h-full object-cover"
              />
            </div>

            <h3 className="mt-4 text-lg font-semibold text-center">{p.name}</h3>

            {p.description && (
              <p className="text-gray-600 text-sm text-center mt-1 px-3 line-clamp-2">
                {p.description}
              </p>
            )}

            <p className="font-bold text-center mt-3 text-gray-900 text-lg">
              {priceFmt.format(p.price ?? 0)}
            </p>

            <button
              onClick={() => agregarAlCarrito(p.id)}
              className="mt-4 w-full bg-black text-white py-2 rounded-md hover:bg-neutral-800 transition"
            >
              Agregar al carrito
            </button>
          </article>
        ))}
      </div>
    </div>
  );
}
