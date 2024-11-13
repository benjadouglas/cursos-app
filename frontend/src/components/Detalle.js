import React from "react";
import { useLocation } from "react-router-dom";
import "./Detalle.css";

const Detalle = () => {
  const location = useLocation();
  const curso = location.state?.curso;

  const handleInscripcion = async () => {
    try {
      const response = await fetch("http://localhost:8085/api/enrollments", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
        body: JSON.stringify({
          id: localStorage.getItem("userId"),
          curso_id: curso.Id,
        }),
      });

      if (!response.ok) {
        throw new Error("Error en la inscripción");
      }

      alert("¡Te has inscrito al curso!");
    } catch (error) {
      console.error("Error:", error);
      alert("Error al inscribirse al curso");
    }
  };

  return (
    <div classname="detalle-container">
      <div classname="detalle-curso">
        <h2>{curso.Nombre}</h2>
        <p>{curso.Id}</p>
        <p>{curso.Capacidad}</p>
        <p>{curso.Precio}</p>
        <button classname="inscripcion-button" onClick={handleInscripcion}>
          inscribirse
        </button>
      </div>
    </div>
  );
};

export default Detalle;
