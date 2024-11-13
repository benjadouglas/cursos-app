import React from 'react';
import './Resultados.css';

const Resultados = ({ cursos }) => {
    return (
        <div>
            <h1>Resultados de la Búsqueda</h1>
            {cursos.length > 0 ? (
                cursos.map(curso => (
                    <div key={curso.id}>
                        <h2>{curso.nombre}</h2>
                        <p>{curso.descripcion}</p>
                        <p>Puntuación: {curso.puntuacion}</p>
                    </div>
                ))
            ) : (
                <p>No se encontraron cursos.</p>
            )}
        </div>
    );
};

export default Resultados;