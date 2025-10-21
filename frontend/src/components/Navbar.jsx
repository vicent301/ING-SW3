import { Link } from "react-router-dom";
import { ShoppingCart, User, Search } from "lucide-react";
import { useState } from "react";

export default function Navbar() {
  const [hovered, setHovered] = useState(null);

  const categories = [
    {
      name: "Calzado",
      items: ["Zapatillas", "Botas", "Sandalias", "Ver Todo"],
    },
    {
      name: "Indumentaria",
      items: ["Remeras", "Buzos", "Pantalones", "Camperas", "Ver Todo"],
    },
    {
      name: "Accesorios",
      items: ["Gorros", "Medias", "Mochilas", "Riñoneras", "Ver Todo"],
    },
  ];

  return (
    <header
      className="sticky top-0 z-50 w-full backdrop-blur-md shadow-md border-b border-neutral-300"
      style={{
        background:
          "linear-gradient(90deg, rgba(255,255,255,0.9) 0%, rgba(230,230,230,0.8) 100%)",
        color: "#111",
      }}
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
                      to={`/productos/${item.toLowerCase()}`}
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
            />
          </div>

          {/* Login */}
          <Link to="/login" className="hover:text-gray-700 transition">
            <User size={24} />
          </Link>

          {/* Carrito */}
          <Link
            to="/carrito"
            className="hover:text-gray-700 transition relative"
          >
            <ShoppingCart size={24} />
            <span className="absolute -top-2 -right-2 bg-gray-700 text-white text-xs font-bold rounded-full px-1.5">
              2
            </span>
          </Link>
        </div>
      </nav>
    </header>
  );
}
