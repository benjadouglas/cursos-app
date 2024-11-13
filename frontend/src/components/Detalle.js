import React from 'react';
import './Detalle.css';

const Detalle = () => {
    const curso = {
        nombre: "Curso de Matemáticas",
        descripcion: "Aprende matemáticas desde cero en este curso completo.",
        puntuacion: 4.5,
    };

    const handleInscripcion = () => {
        alert("¡Te has inscrito al curso!");
    };

    return (
        <div className="detalle-container">
            <div className="detalle-curso">
                <h2>{curso.nombre}</h2>
                <p>{curso.descripcion}</p>
                <div className="puntuacion">Puntuación: {curso.puntuacion}</div>
                <button className="inscripcion-button" onClick={handleInscripcion}>Inscribirse</button>
            </div>
        </div>
    );
};

export default Detalle;
