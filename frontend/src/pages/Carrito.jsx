import { ShoppingCart } from "lucide-react";

export default function Carrito() {
  return (
    <section className="flex flex-col items-center justify-center text-center py-16 px-6">
      <h1 className="text-4xl font-extrabold text-white mb-3 tracking-tight flex items-center gap-2">
        <ShoppingCart size={32} /> Tu <span className="text-green-500">Carrito</span>
      </h1>
      <p className="text-gray-400 max-w-md">
        Aún no agregaste productos. Navegá por el catálogo y empezá a armar tu colección 👟
      </p>
    </section>
  );
}
