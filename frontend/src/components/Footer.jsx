export default function Footer() {
  return (
    <footer className="bg-neutral-950 text-gray-400 py-8 border-t border-neutral-800 text-center">
      <p className="text-sm">
        © 2025 <span className="text-green-500 font-semibold">ZapasStore</span> — Todos los derechos reservados
      </p>
      <div className="flex justify-center gap-6 mt-3 text-xl">
        <a href="#" className="hover:text-green-500 transition">🐦</a>
        <a href="#" className="hover:text-green-500 transition">📸</a>
        <a href="#" className="hover:text-green-500 transition">💬</a>
      </div>
    </footer>
  );
}
