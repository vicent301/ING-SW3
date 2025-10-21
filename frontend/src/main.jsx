import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import "./index.css";
import Home from "./pages/Home";
import Productos from "./pages/Productos";
import Carrito from "./pages/Carrito";
import Navbar from "./components/Navbar";

function App() {
  return (
    <BrowserRouter>
      {/* Contenedor principal */}
      <div className="min-h-screen flex flex-col bg-neutral-950 text-white font-sans overflow-x-hidden">
        <Navbar />

        {/* Contenido centrado */}
        <main className="flex-grow flex justify-center items-start w-full px-6 md:px-12 lg:px-20 py-12">
          <div className="w-full max-w-7xl mx-auto">
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/productos" element={<Productos />} />
              <Route path="/carrito" element={<Carrito />} />
            </Routes>
          </div>
        </main>

        {/* Footer */}
        <footer className="text-center text-gray-500 text-sm py-6 border-t border-neutral-800 mt-auto">
          © 2025 ZapasStore — Todos los derechos reservados
        </footer>
      </div>
    </BrowserRouter>
  );
}

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
