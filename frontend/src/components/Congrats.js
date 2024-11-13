import React from 'react';

const Congrats = ({ success, mensaje }) => {
    return (
        <div>
            {success ? <h2>¡Inscripción exitosa!</h2> : <h2>Error en la inscripción</h2>}
            <p>{mensaje}</p>
        </div>
    );
};

export default Congrats;
