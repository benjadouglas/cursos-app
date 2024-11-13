import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import "./Cursos.css";

const Cursos = () => {
  const navigate = useNavigate();
  const [cursos, setCursos] = useState([]);

  useEffect(() => {
    const fetchCursos = async () => {
      try {
        const response = await fetch(
          "http://localhost:8080/search?q=Id:*&offset=0&limit=10000",
          {
            method: "GET",
            mode: "cors",
            headers: {
              "Content-Type": "application/json",
            },
          },
        );

        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        setCursos(data);
      } catch (error) {
        console.error("Error fetching courses:", error);
        setCursos([]);
      }
    };

    fetchCursos();
  }, []);

  const goToDetalle = (curso) => {
    navigate("/detalle", { state: { curso } });
  };

  return (
    <div className="cursos-lista">
      {cursos.map((curso) => (
        <div key={curso.id} className="curso-item">
          <h3>{curso.Nombre}</h3>
          <p>{curso.Precio}</p>
          <p>Capacidad: {curso.Capacidad}</p>
          <button onClick={() => goToDetalle(curso)} className="detalle-button">
            Ver Detalle
          </button>
        </div>
      ))}
    </div>
  );
};

export default Cursos;
