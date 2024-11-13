import React, { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import "./Home.css";

const Home = ({ setCursos }) => {
  const [search, setSearch] = useState("");
  const navigate = useNavigate();

  console.log(search);
  const handleSearch = async () => {
    try {
      // Codificar los parámetros de búsqueda para manejar caracteres especiales

      const response = await fetch(
        `http://localhost:8080/search?q=Nombre:${search}~&offset=0&limit=10000`,
        {
          method: "GET",
          mode: "cors",
          headers: {
            "Content-Type": "application/json",
          },
        },
      );

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || `Error HTTP: ${response.status}`);
      }

      const data = await response.json();

      setCursos(data);
      navigate("/resultados");
    } catch (error) {
      console.error("Error en la búsqueda:", error);
      setCursos([]);
    }
  };

  return (
    <div className="home-container">
      <h1 className="home-title">Bienvenido a la Plataforma de Cursos</h1>
      <p className="home-subtitle">Encuentra y aprende a tu ritmo</p>

      <div className="home-search">
        <input
          type="text"
          placeholder="Buscar cursos..."
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          className="search-input"
        />
        <button onClick={handleSearch} className="search-button">
          Buscar
        </button>
      </div>

      <div className="home-buttons">
        <Link to="/registro" className="home-link">
          <button className="action-button">Registrarse</button>
        </Link>
        <Link to="/login" className="home-link">
          <button className="action-button">Iniciar Sesión</button>
        </Link>
        <Link to="/cursos" className="home-link">
          <button className="action-button">Ver Cursos</button>
        </Link>
      </div>

      <div className="home-bottom-image"></div>
    </div>
  );
};

export default Home;
