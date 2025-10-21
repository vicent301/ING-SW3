import { useEffect, useState } from "react";
import logo from "../assets/logo.png";

export default function Productos() {
  const [productos, setProductos] = useState([]);

  useEffect(() => {
    setProductos([
      {
        id: 1,
        name: "Nike SB Dunk Low",
        description: "Zapatilla skate clásica con diseño moderno.",
        price: 120,
        image_url: "https://via.placeholder.com/250x250?text=Nike+SB",
      },
      {
        id: 2,
        name: "Adidas Forum Low",
        description: "Inspiración retro, comodidad urbana.",
        price: 140,
        image_url: "https://via.placeholder.com/250x250?text=Adidas+Forum",
      },
      {
        id: 3,
        name: "Vans Old Skool",
        description: "Diseño atemporal con suela waffle.",
        price: 110,
        image_url: "https://via.placeholder.com/250x250?text=Vans+Old+Skool",
      },
      {
        id: 4,
        name: "Puma Suede Classic",
        description: "Comodidad y estilo en cada paso.",
        price: 100,
        image_url: "https://via.placeholder.com/250x250?text=Puma+Suede",
      },
      {
        id: 5,
        name: "Nike Air Max",
        description: "Estilo urbano con máxima comodidad.",
        price: 150,
        image_url: "https://via.placeholder.com/250x250?text=Nike+Air+Max",
      },
      {
        id: 6,
        name: "Converse All Star",
        description: "Clásico atemporal que nunca falla.",
        price: 90,
        image_url: "https://via.placeholder.com/250x250?text=Converse",
      },
    ]);
  }, []);

  return (
    <div style={{ backgroundColor: "white", color: "#111", minHeight: "100vh" }}>
      {/* 🔹 NAVBAR */}
      <header
        style={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
          padding: "1rem 2rem",
          borderBottom: "1px solid #ddd",
        }}
      >
        {/* Buscador */}
        <div style={{ display: "flex", alignItems: "center", width: "30%" }}>
          <input
            type="text"
            placeholder="Buscar productos..."
            style={{
              width: "100%",
              padding: "0.5rem 0.75rem",
              border: "1px solid #ccc",
              borderRadius: "4px 0 0 4px",
              outline: "none",
            }}
          />
          <button
            style={{
              padding: "0.5rem 0.75rem",
              border: "1px solid #ccc",
              borderLeft: "none",
              backgroundColor: "#f3f3f3",
              borderRadius: "0 4px 4px 0",
              cursor: "pointer",
            }}
          >
            🔍
          </button>
        </div>

        {/* Logo */}
        <div style={{ flex: 1, textAlign: "center" }}>
          <img
            src={logo}
            alt="ZapaStore"
            style={{
              height: "50px",
              objectFit: "contain",
            }}
          />
        </div>

        {/* Espacio vacío a la derecha */}
        <div style={{ width: "30%" }}></div>
      </header>

      {/* 🔹 CATÁLOGO */}
      <main style={{ padding: "2rem" }}>
        <h2
          style={{
            textAlign: "center",
            fontSize: "2rem",
            fontWeight: "bold",
            marginBottom: "2rem",
          }}
        >
          Catálogo de Productos
        </h2>

        {/* 💥 GRILLA DE 3 COLUMNAS */}
        <div
          style={{
            display: "grid",
            gridTemplateColumns: "repeat(3, 1fr)",
            gap: "2.5rem",
            justifyItems: "center",
            maxWidth: "1200px",
            margin: "0 auto",
          }}
        >
          {productos.map((p) => (
            <div
              key={p.id}
              style={{
                backgroundColor: "white",
                border: "1px solid #ddd",
                borderRadius: "12px",
                boxShadow: "0 2px 6px rgba(0,0,0,0.1)",
                transition: "transform 0.3s, box-shadow 0.3s",
                width: "250px",
                textAlign: "center",
              }}
              onMouseEnter={(e) => {
                e.currentTarget.style.transform = "scale(1.03)";
                e.currentTarget.style.boxShadow = "0 4px 12px rgba(0,0,0,0.15)";
              }}
              onMouseLeave={(e) => {
                e.currentTarget.style.transform = "scale(1)";
                e.currentTarget.style.boxShadow = "0 2px 6px rgba(0,0,0,0.1)";
              }}
            >
              <div
                style={{
                  height: "220px",
                  backgroundColor: "#f3f3f3",
                  borderRadius: "12px 12px 0 0",
                  overflow: "hidden",
                }}
              >
                <img
                  src={p.image_url}
                  alt={p.name}
                  style={{
                    width: "100%",
                    height: "100%",
                    objectFit: "cover",
                  }}
                />
              </div>
              <div style={{ padding: "1rem" }}>
                <h3 style={{ fontSize: "1.1rem", fontWeight: "600" }}>
                  {p.name}
                </h3>
                <p
                  style={{
                    color: "#666",
                    fontSize: "0.9rem",
                    margin: "0.5rem 0",
                  }}
                >
                  {p.description}
                </p>
                <p style={{ fontWeight: "bold", marginBottom: "0.75rem" }}>
                  ${p.price}
                </p>
                <button
                  style={{
                    width: "100%",
                    padding: "0.5rem",
                    backgroundColor: "black",
                    color: "white",
                    border: "none",
                    borderRadius: "6px",
                    cursor: "pointer",
                  }}
                >
                  Agregar al carrito
                </button>
              </div>
            </div>
          ))}
        </div>
      </main>

      {/* 🔹 FOOTER */}
      <footer
        style={{
          textAlign: "center",
          padding: "1.5rem",
          borderTop: "1px solid #ddd",
          fontSize: "0.9rem",
          color: "#666",
        }}
      >
        © {new Date().getFullYear()} ZapaStore — Todos los derechos reservados.
      </footer>
    </div>
  );
}
