// src/pages/Productos.jsx
import { useEffect, useMemo, useState } from "react";
import { useLocation } from "react-router-dom";
import logo from "../assets/logo.png";
import toast from "react-hot-toast";
import { getProducts, addToCart } from "../services/api";

export default function Productos() {
  const [productos, setProductos] = useState([]);
  const [loading, setLoading] = useState(true);
  const [errMsg, setErrMsg] = useState("");
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

  // ✅ cargar productos desde services/api (respeta VITE_API_URL)
  useEffect(() => {
    let mounted = true;
    async function cargar() {
      try {
        const data = await getProducts(); // <- usa services/api
        if (!mounted) return;
        setProductos(Array.isArray(data) ? data : []);
      } catch (e) {
        console.error(e);
        setErrMsg("No se pudieron cargar los productos.");
      } finally {
        if (mounted) setLoading(false);
      }
    }
    cargar();
    return () => (mounted = false);
  }, []);

  // ✅ aplicar filtro del buscador
  const productosFiltrados = productos.filter(
    (p) =>
      p.name?.toLowerCase().includes(search) ||
      p.description?.toLowerCase().includes(search)
  );

  // ✅ agregar al carrito (usa services/api; maneja token adentro)
  async function agregarAlCarrito(productId) {
    try {
      await addToCart(productId, 1);
      toast.success("✅ Producto agregado al carrito");
    } catch (e) {
      // Si no hay token o 401, tu services/api ya lanza error
      toast.error("Iniciá sesión para agregar al carrito");
    }
  }

  return (
    <div className="w-full" data-testid="products-page">
      {/* ✅ Logo (chico y cercano al navbar) */}
      <div className="flex justify-center mt-6 mb-2">
        <img src={logo} alt="ZapaStore logo" className="w-20 opacity-90" />
      </div>

      {/* ✅ Título */}
      <h2 className="text-center text-3xl font-bold mt-1 mb-8 tracking-wide">
        Catálogo de Productos
      </h2>

      {/* ✅ Estados */}
      {loading && (
        <div className="min-h-[30vh] flex items-center justify-center">
          <p className="text-gray-500" data-testid="products-loading">
            Cargando productos…
          </p>
        </div>
      )}

      {!loading && errMsg && (
        <div className="min-h-[30vh] flex items-center justify-center">
          <p className="text-red-600" data-testid="products-error">
            {errMsg}
          </p>
        </div>
      )}

      {!loading && !errMsg && search && productosFiltrados.length === 0 && (
        <p className="text-center text-gray-500 mb-8" data-testid="no-results">
          No se encontraron resultados para{" "}
          <span className="font-semibold">“{params.get("search")}”</span>.
        </p>
      )}

      {/* ✅ Grid (1 / 2 / 3 columnas) */}
      {!loading && !errMsg && (
        <div
          className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-10 max-w-7xl mx-auto px-6 pb-12"
          data-testid="products-grid"
        >
          {productosFiltrados.map((p) => (
            <article
              key={p.id}
              data-testid="product-card"
              className="bg-white border border-gray-200 rounded-xl shadow-md hover:shadow-xl transition p-4"
            >
              <div className="w-full h-56 bg-gray-100 rounded-lg overflow-hidden">
                <img
                  src={p.image_url || "https://via.placeholder.com/600x400?text=Producto"}
                  alt={p.name}
                  className="w-full h-full object-cover"
                />
              </div>

              <h3 className="mt-4 text-lg font-semibold text-center">
                {p.name}
              </h3>

              {p.description && (
                <p className="text-gray-600 text-sm text-center mt-1 px-3 line-clamp-2">
                  {p.description}
                </p>
              )}

              <p className="font-bold text-center mt-3 text-gray-900 text-lg">
                {priceFmt.format(p.price ?? 0)}
              </p>

              <button
                data-testid="add-to-cart"
                onClick={() => agregarAlCarrito(p.id)}
                className="mt-4 w-full bg-black text-white py-2 rounded-md hover:bg-neutral-800 transition"
              >
                Agregar al carrito
              </button>
            </article>
          ))}
        </div>
      )}
    </div>
  );
}
