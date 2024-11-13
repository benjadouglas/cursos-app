import React from 'react';

const MisCursos = ({ cursos }) => {
    return (
        <div>
            <h2>Mis Cursos</h2>
            {cursos.map((curso) => (
                <div key={curso.id}>
                    <h3>{curso.titulo}</h3>
                    <p>{curso.descripcion}</p>
                </div>
            ))}
        </div>
    );
};

export default MisCursos;
