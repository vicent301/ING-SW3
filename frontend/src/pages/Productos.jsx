export default function Productos() {
  return (
    <section className="w-full min-h-screen bg-gradient-to-b from-neutral-950 to-neutral-900 flex flex-col items-center py-20 px-6">
      <div className="max-w-7xl w-full text-center">
        {/* 🏷️ Título */}
        <h1 className="text-6xl font-extrabold text-white mb-6 tracking-tight drop-shadow-[0_2px_8px_rgba(0,255,100,0.3)]">
          Catálogo de <span className="text-green-500">Productos</span>
        </h1>

        <p className="text-gray-400 max-w-2xl mx-auto mb-16 text-lg leading-relaxed">
          Descubrí nuestras zapatillas más populares y encontrá tu estilo 👟  
          <span className="block mt-2 text-sm text-gray-500">
            (Los productos se cargarán automáticamente desde la base de datos)
          </span>
        </p>

        {/* 🧱 Grilla FIJA de 4 columnas con más separación */}
<div className="grid grid-cols-4 justify-items-center w-full gap-x-10 gap-y-[6rem] mt-10">
  {[1, 2, 3, 4, 5, 6, 7, 8].map((p) => (
    <div
      key={p}
      className="group bg-neutral-900 border border-neutral-800 rounded-2xl p-5 w-64 hover:border-green-500 hover:shadow-[0_0_20px_rgba(0,255,100,0.15)] hover:-translate-y-2 transition-all duration-300 ease-in-out"
    >
      {/* Imagen del producto */}
      <div className="h-48 bg-neutral-800 rounded-xl flex items-center justify-center overflow-hidden">
        <img
          src={`https://via.placeholder.com/200x200?text=Zapa+${p}`}
          alt={`Zapatilla ${p}`}
          className="opacity-80 group-hover:opacity-100 group-hover:scale-110 transition-all duration-500 ease-in-out"
        />
      </div>

      {/* Info del producto */}
      <h3 className="text-lg font-semibold mt-4 text-white group-hover:text-green-400 transition">
        Zapatilla #{p}
      </h3>
      <p className="text-green-400 font-bold mt-1 text-lg">$120</p>

      {/* Botón */}
      <button className="w-full bg-green-500 text-black py-2 rounded-full mt-4 font-semibold hover:bg-green-400 transition-colors shadow-md shadow-green-500/20">
        Agregar al carrito
      </button>
    </div>
  ))}
</div>

      </div>
    </section>
  );
}
