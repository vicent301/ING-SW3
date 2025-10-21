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
    <header className="bg-black text-white shadow-md border-b border-neutral-800 sticky top-0 z-50 w-full">
      <nav className="max-w-[1600px] mx-auto px-16 py-5 flex justify-between items-center">

        {/* Logo */}
        <Link
          to="/"
          className="text-4xl font-extrabold tracking-tight hover:text-green-400 transition"
        >
          <span className="text-white">Zapas</span>
          <span className="text-green-500">Store</span>
        </Link>

        {/* Categories */}
        <ul className="hidden md:flex gap-12 text-base font-semibold uppercase tracking-wide">
          {categories.map((cat) => (
            <li
              key={cat.name}
              onMouseEnter={() => setHovered(cat.name)}
              onMouseLeave={() => setHovered(null)}
              className="relative cursor-pointer hover:text-green-400 transition"
            >
              {cat.name}
              {hovered === cat.name && (
                <div className="absolute left-0 top-full mt-2 bg-neutral-900 text-white rounded-md shadow-lg py-4 px-6 grid grid-cols-1 gap-2 w-48 border border-neutral-700 animate-fadeIn">
                  {cat.items.map((item) => (
                    <Link
                      key={item}
                      to={`/productos/${item.toLowerCase()}`}
                      className="hover:text-green-400 transition text-sm"
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
          <div className="hidden md:flex items-center bg-neutral-800 rounded-full px-4 py-2">
            <Search size={18} className="text-gray-400" />
            <input
              type="text"
              placeholder="Buscar..."
              className="bg-transparent text-sm text-gray-300 px-2 focus:outline-none w-28"
            />
          </div>

          {/* Login */}
          <Link to="/login" className="hover:text-green-400 transition">
            <User size={24} />
          </Link>

          {/* Carrito */}
          <Link to="/carrito" className="hover:text-green-400 transition relative">
            <ShoppingCart size={24} />
            <span className="absolute -top-2 -right-2 bg-green-500 text-black text-xs font-bold rounded-full px-1.5">
              2
            </span>
          </Link>
        </div>
      </nav>
    </header>
  );
}
