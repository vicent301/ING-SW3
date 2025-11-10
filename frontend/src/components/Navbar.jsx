import { Link, useNavigate } from "react-router-dom";
import { ShoppingCart, User, Search, LogOut } from "lucide-react";
import { useState, useEffect } from "react";
import { getProfile, getCart } from "../services/api";

export default function Navbar() {
  const [user, setUser] = useState(null);
  const [cartCount, setCartCount] = useState(0);
  const [search, setSearch] = useState("");
  const navigate = useNavigate();

  const categories = [
    { name: "Calzado", link: "/products" },
    { name: "Indumentaria", link: "/products" },
  ];

  useEffect(() => {
    async function loadData() {
      try {
        const userData = await getProfile();
        setUser(userData);
        const cartData = await getCart();
        setCartCount(cartData.items?.length || 0);
      } catch {}
    }
    loadData();
  }, []);

  const handleSearch = (e) => {
    if (e.key === "Enter") {
      navigate(`/products?search=${encodeURIComponent(search)}`);
    }
  };

  const handleLogout = () => {
    localStorage.removeItem("token");
    setUser(null);
    navigate("/login");
  };

  return (
    <header
      className="sticky top-0 z-50 w-full backdrop-blur-md shadow-md border-b border-neutral-300"
      style={{
        background:
          "linear-gradient(90deg, rgba(255,255,255,0.9) 0%, rgba(230,230,230,0.8) 100%)",
        color: "#111",
      }}
    >
      <nav className="max-w-[1440px] mx-auto px-4 md:px-8 lg:px-10 py-4 flex items-center gap-4">
        {/* Columna 1: Logo */}
        <div className="min-w-[140px] flex">
          <Link
            to="/"
            className="text-3xl md:text-4xl font-extrabold tracking-tight transition hover:opacity-90"
            style={{
              background:
                "linear-gradient(90deg, #000000 0%, #333333 50%, #666666 100%)",
              WebkitBackgroundClip: "text",
              WebkitTextFillColor: "transparent",
              fontFamily: "'Anton', sans-serif",
              letterSpacing: "1px",
            }}
          >
            ZapaStore
          </Link>
        </div>

        {/* Columna 2: Categorías centradas */}
        <ul className="flex-1 flex justify-center gap-8 md:gap-12 text-sm md:text-base font-semibold uppercase tracking-wide">
          {categories.map((cat) => (
            <li key={cat.name}>
              <Link
                to={cat.link}
                className="hover:text-gray-700 transition relative after:content-[''] after:absolute after:left-0 after:-bottom-1 after:h-[2px] after:w-0 hover:after:w-full after:bg-gray-800 after:transition-all"
              >
                {cat.name}
              </Link>
            </li>
          ))}
        </ul>

        {/* Columna 3: Buscador + iconos */}
        <div className="min-w-[210px] flex items-center justify-end gap-4">
          <div className="hidden md:flex items-center bg-white/50 rounded-full px-3 py-2 shadow-sm backdrop-blur">
            <Search size={16} className="text-gray-700" />
            <input
              type="text"
              placeholder="Buscar..."
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              onKeyDown={handleSearch}
              className="bg-transparent text-sm text-gray-800 px-2 focus:outline-none w-28"
            />
          </div>

          {user ? (
            <div className="hidden sm:flex items-center gap-3">
              <span className="font-semibold text-sm">Hola, {user.name}</span>
              <button onClick={handleLogout} title="Cerrar sesión">
                <LogOut size={20} className="text-gray-700 hover:text-black transition" />
              </button>
            </div>
          ) : (
            <Link to="/login" className="hover:text-gray-700 transition">
              <User size={22} />
            </Link>
          )}

          <Link to="/carrito" className="hover:text-gray-700 transition relative">
            <ShoppingCart size={22} />
            {cartCount > 0 && (
              <span className="absolute -top-2 -right-2 bg-gray-800 text-white text-xs font-bold rounded-full px-1.5">
                {cartCount}
              </span>
            )}
          </Link>
        </div>
      </nav>
    </header>
  );
}
