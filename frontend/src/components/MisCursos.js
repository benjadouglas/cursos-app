import React, { useState, useEffect } from "react";
import "./Cursos.css";

const MisCursos = () => {
  const [cursos, setCursos] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchCursos = async () => {
      try {
        // Obtener el token del localStorage
        const token = localStorage.getItem("token");
        const userId = localStorage.getItem("userId"); // Asegúrate de guardar el userId durante el login

        if (!token || !userId) {
          throw new Error("No hay sesión activa");
        }

        // 1. Primero obtener las inscripciones del usuario
        const enrollmentsResponse = await fetch(
          `http://localhost:8085/api/enrollments/user/${userId}`,
          {
            method: "GET",
            mode: "cors",
            headers: {
              Authorization: `Bearer ${token}`,
              "Content-Type": "application/json",
            },
          },
        );

        if (!enrollmentsResponse.ok) {
          throw new Error("Error al obtener las inscripciones");
        }

        const enrollments = await enrollmentsResponse.json();
        console.log(enrollments);

        setCursos(enrollments);
      } catch (err) {
        setError(err.message);
        console.error("Error:", err);
      } finally {
        setLoading(false);
      }
    };

    fetchCursos();
  }, []);

  if (loading) {
    return <div className="loading">Cargando cursos...</div>;
  }

  if (error) {
    return <div className="error">Error: {error}</div>;
  }

  return (
    <div className="mis-cursos-container">
      <h2 className="mis-cursos-title">Mis Cursos</h2>
      {cursos.length === 0 ? (
        <p className="no-cursos">No estás inscrito en ningún curso todavía.</p>
      ) : (
        <div className="cursos-grid">
          {cursos.map((curso) => (
            <div key={curso._id} className="curso-card">
              <div className="curso-info">
                <h3 className="curso-titulo">{curso.Nombre}</h3>
                <p className="curso-descripcion">{curso.Precio}</p>
                <p>{curso.Capacidad}</p>
                <p classname="curso-profesor">{curso.Profesor}</p>
                <button
                  className="ver-curso-btn"
                  onClick={() => (window.location.href = `/curso/${curso._id}`)}
                >
                  Ver Curso
                </button>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default MisCursos;
