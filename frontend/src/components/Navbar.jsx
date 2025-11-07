import { Link, useNavigate } from "react-router-dom";
import { ShoppingCart, User, Search, LogOut } from "lucide-react";
import { useState, useEffect } from "react";
import { getProfile, getCart } from "../services/api";

export default function Navbar() {
  const [hovered, setHovered] = useState(null);
  const [user, setUser] = useState(null);
  const [cartCount, setCartCount] = useState(0);
  const navigate = useNavigate();

  const categories = [
    { name: "Calzado", items: ["Zapatillas", "Botas", "Sandalias", "Ver Todo"] },
    { name: "Indumentaria", items: ["Remeras", "Buzos", "Pantalones", "Camperas", "Ver Todo"] },
    { name: "Accesorios", items: ["Gorros", "Medias", "Mochilas", "Riñoneras", "Ver Todo"] },
  ];

  useEffect(() => {
    async function loadData() {
      try {
        const userData = await getProfile();
        setUser(userData);

        const cartData = await getCart();
        setCartCount(cartData.items?.length || 0);
      } catch (err) {
        // sin token o 401 -> ignorar
      }
    }
    loadData();
  }, []);

  const handleLogout = () => {
    localStorage.removeItem("token");
    setUser(null);
    navigate("/login");
  };

  return (
    <header
      className="sticky top-0 z-50 w-full backdrop-blur-md shadow-md border-b border-neutral-300"
      style={{
        background: "linear-gradient(90deg, rgba(255,255,255,0.9) 0%, rgba(230,230,230,0.8) 100%)",
        color: "#111",
      }}
      data-testid="navbar"
    >
      <nav className="max-w-[1600px] mx-auto px-16 py-5 flex justify-between items-center">
        {/* Logo */}
        <Link
          to="/"
          className="text-4xl font-extrabold tracking-tight transition duration-300 hover:opacity-90"
          style={{
            background: "linear-gradient(90deg, #000000 0%, #333333 50%, #666666 100%)",
            WebkitBackgroundClip: "text",
            WebkitTextFillColor: "transparent",
            fontFamily: "'Anton', sans-serif",
            letterSpacing: "1px",
          }}
          data-testid="logo-home"
        >
          ZapaStore
        </Link>

        {/* Categorías */}
        <ul className="hidden md:flex gap-12 text-base font-semibold uppercase tracking-wide">
          {categories.map((cat) => (
            <li
              key={cat.name}
              onMouseEnter={() => setHovered(cat.name)}
              onMouseLeave={() => setHovered(null)}
              className="relative cursor-pointer hover:text-gray-700 transition"
            >
              {cat.name}
              {hovered === cat.name && (
                <div className="absolute left-0 top-full mt-2 bg-white/80 text-gray-900 rounded-md shadow-lg py-4 px-6 grid grid-cols-1 gap-2 w-48 border border-gray-300 animate-fadeIn backdrop-blur-md">
                  {cat.items.map((item) => (
                    <Link
                      key={item}
                      to={`/products/${item.toLowerCase()}`}
                      className="hover:text-gray-600 transition text-sm"
                    >
                      {item}
                    </Link>
                  ))}
                </div>
              )}
            </li>
          ))}
        </ul>

        {/* Search + Icons */}
        <div className="flex items-center gap-6 pr-6">
          {/* Search */}
          <div className="hidden md:flex items-center bg-white/40 rounded-full px-4 py-2 shadow-sm backdrop-blur-md">
            <Search size={18} className="text-gray-700" />
            <input
              type="text"
              placeholder="Buscar..."
              className="bg-transparent text-sm text-gray-800 px-2 focus:outline-none w-28"
              data-testid="search-input"
            />
          </div>

          {/* Usuario */}
          {user ? (
            <div className="flex items-center gap-4">
              <span className="font-semibold text-sm" data-testid="greeting">Hola, {user.name}</span>
              <button onClick={handleLogout} title="Cerrar sesión" data-testid="logout-button">
                <LogOut size={22} className="text-gray-700 hover:text-black transition" />
              </button>
            </div>
          ) : (
            <Link to="/login" className="hover:text-gray-700 transition" data-testid="nav-login">
              <User size={24} />
            </Link>
          )}

          {/* Carrito */}
          <Link to="/carrito" className="hover:text-gray-700 transition relative" data-testid="navbar-cart">
            <ShoppingCart size={24} />
            {cartCount > 0 && (
              <span className="absolute -top-2 -right-2 bg-gray-700 text-white text-xs font-bold rounded-full px-1.5" data-testid="cart-badge">
                {cartCount}
              </span>
            )}
          </Link>
        </div>
      </nav>
    </header>
  );
}
