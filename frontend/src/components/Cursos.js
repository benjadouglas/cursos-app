import React from 'react';
import { useNavigate } from 'react-router-dom';
import './Cursos.css';

const Cursos = () => {
    const navigate = useNavigate();

    const cursos = [
        { id: 1, nombre: "Curso de Matemáticas", descripcion: "Aprende matemáticas desde cero", puntuacion: 4.5 },
        { id: 2, nombre: "Curso de Historia", descripcion: "Historia universal", puntuacion: 4.2 },
        { id: 3, nombre: "Curso de Ciencias", descripcion: "Explora el mundo de la ciencia", puntuacion: 4.7 },
    ];

    const goToDetalle = (curso) => {
        navigate('/detalle', { state: { curso } });
    };

    return (
        <div className="cursos-lista">
            {cursos.map((curso) => (
                <div key={curso.id} className="curso-item">
                    <h3>{curso.nombre}</h3>
                    <p>{curso.descripcion}</p>
                    <p>Puntuación: {curso.puntuacion}</p>
                    <button onClick={() => goToDetalle(curso)} className="detalle-button">Ver Detalle</button>
                </div>
            ))}
        </div>
    );
};

export default Cursos;