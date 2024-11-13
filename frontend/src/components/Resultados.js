import React from "react";
import "./Resultados.css";

const Resultados = ({ cursos }) => {
  const handleInscripcion = async (cursoId) => {
    try {
      const response = await fetch("http://localhost:8085/api/enrollments", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
        body: JSON.stringify({
          id: localStorage.getItem("userId"),
          curso_id: cursoId,
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
    <div className="resultados-container">
      <h1>Resultados de la Búsqueda</h1>
      {cursos.length > 0 ? (
        cursos.map((curso) => (
          <div key={curso.id} className="resultado-item">
            <h2>{curso.Nombre}</h2>
            <p>Precio: {curso.Precio}</p>
            <p>Capacidad: {curso.capacidad}</p>
            <p>Profesor: {curso.Profesor}</p>
            <p>Duracion: {curso.duracion}</p>
            <button onClick={() => handleInscripcion(curso.Id)}>
              Inscribirse
            </button>
          </div>
        ))
      ) : (
        <p className="no-resultados">No se encontraron cursos.</p>
      )}
    </div>
  );
};

export default Resultados;
