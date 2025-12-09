import { Link } from "react-router-dom";
import bg from "../assets/skater-bg.png"; // tu imagen

export default function Home() {
  return (
    <div
      className="relative flex flex-col items-center justify-center min-h-screen text-center overflow-hidden"
      style={{
        backgroundImage: `url(${bg})`,
        backgroundSize: "cover",
        backgroundPosition: "center",
        backgroundRepeat: "no-repeat",
      }}
    >
      {/* 🔹 Capa oscura translúcida real */}
      <div
        style={{
          position: "absolute",
          inset: 0,
          backgroundColor: "rgba(0, 0, 0, 0.65)", // <- opacidad fuerte
          backdropFilter: "blur(2px)", // <- desenfoque sutil
          zIndex: 0,
        }}
      ></div>

      {/* 🔹 Contenido visible */}
      <div className="relative z-10 max-w-3xl mx-auto px-6 text-center"
     style={{ color: "rgb(40,20,120)" }}> 

        <h1
          className="text-6xl md:text-7xl font-bold mb-6 drop-shadow-lg"
          style={{
            fontFamily: "'Anton', sans-serif",
            letterSpacing: "2px",
          }}
        >
          Bienvenido a ZapaStore 
          
        </h1>

        <p
          className="text-lg md:text-xl mb-8 text-gray-200"
          style={{
            fontFamily: "'Oswald', sans-serif",
            fontWeight: "400",
          }}
        >
          Encontrá tus zapatillas favoritas al mejor precio y en un solo lugar.
          anda?
        </p>

        <Link
          to="/products"
          className="inline-block bg-white text-black px-8 py-3 rounded-md font-semibold hover:bg-gray-200 transition"
          style={{
            fontFamily: "'Oswald', sans-serif",
            letterSpacing: "1px",
          }}
        >
          Ver Catálogo
        </Link>
      </div>
    </div>
  );
}
