import React, { useState } from 'react';
import './Login.css';

const Login = ({ onLogin }) => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const handleLogin = () => {
        onLogin(username, password);
    };

    return (
        <div className="login-container">
            <h2 className="login-title">Iniciar Sesión</h2>
            <input
                type="text"
                placeholder="Usuario"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                className="login-input"
            />
            <input
                type="password"
                placeholder="Contraseña"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="login-input"
            />
            <button onClick={handleLogin} className="login-button">Ingresar</button>
        </div>
    );
};

export default Login;
