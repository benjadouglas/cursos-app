import React, { useState } from 'react';
import './Registro.css';

const Register = ({ onRegister }) => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const handleRegister = () => {
        onRegister(username, password);
    };

    return (
        <div className="register-container">
            <h2 className="register-title">Registro</h2>
            <input
                type="text"
                placeholder="Usuario"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                className="register-input"
            />
            <input
                type="password"
                placeholder="Contraseña"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="register-input"
            />
            <button onClick={handleRegister} className="register-button">Registrarse</button>
        </div>
    );
};

export default Register;
