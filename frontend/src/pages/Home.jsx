import { Link } from "react-router-dom";

export default function Home() {
  return (
    <section
      className="h-[90vh] bg-cover bg-center flex items-center justify-center text-center"
      style={{
        backgroundImage:
          "url('https://images.unsplash.com/photo-1528701800489-20be3c9671f0?auto=format&fit=crop&w=1920&q=80')",
      }}
    >
      <div className="bg-black bg-opacity-60 p-12 rounded-3xl max-w-2xl">
        <h1 className="text-5xl font-extrabold mb-4 text-white">
          Bienvenido a <span className="text-green-500">ZapasStore</span> 👟
        </h1>
        <p className="text-gray-300 mb-8 text-lg">
          Encontrá tus zapatillas favoritas al mejor precio y en un solo lugar.
        </p>
        <Link
          to="/productos"
          className="bg-green-500 text-black px-8 py-3 rounded-full font-semibold text-lg hover:bg-green-400 transition"
        >
          Ver Catálogo
        </Link>
      </div>
    </section>
  );
}
