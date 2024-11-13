import React from "react";
import "./Resultados.css";

const Resultados = ({ cursos }) => {
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
          </div>
        ))
      ) : (
        <p className="no-resultados">No se encontraron cursos.</p>
      )}
    </div>
  );
};

export default Resultados;
