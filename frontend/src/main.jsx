import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route, useLocation } from "react-router-dom";
import "./index.css";

// 🧩 Páginas
import Home from "./pages/Home";
import Productos from "./pages/Productos";
import Carrito from "./pages/Carrito";
import Login from "./pages/Login";
import Register from "./pages/Register";

// 🧱 Componentes
import Navbar from "./components/Navbar";

function Layout() {
  const location = useLocation();
  // Ocultamos navbar y footer en login/register
  const hideLayout = ["/login", "/register"].includes(location.pathname);

  return (
    <div className="min-h-screen flex flex-col bg-white text-gray-900 font-sans overflow-x-hidden">
      {!hideLayout && <Navbar />}

      {/* Contenido principal */}
      <main className="flex-grow flex justify-center items-start w-full px-6 md:px-12 lg:px-20 py-12">
        <div className="w-full max-w-7xl mx-auto">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/products" element={<Productos />} />
            <Route path="/carrito" element={<Carrito />} />
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
          </Routes>
        </div>
      </main>

      {!hideLayout && (
        <footer className="text-center text-gray-500 text-sm py-6 border-t border-neutral-800 mt-auto">
          © 2025 ZapaStore — Todos los derechos reservados
        </footer>
      )}
    </div>
  );
}

function App() {
  return (
    <BrowserRouter>
      <Layout />
    </BrowserRouter>
  );
}

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
