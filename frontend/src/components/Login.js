import React, { useState } from "react";
import "./Login.css";

const Login = ({ onLogin }) => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  const handleLogin = async () => {
    try {
      // Validación básica
      if (!username || !password) {
        setError("Por favor, complete todos los campos");
        return;
      }

      console.log("Intentando iniciar sesión...");
      const response = await fetch("http://localhost:8085/api/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        mode: "cors",
        body: JSON.stringify({
          email: username,
          password: password,
        }),
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.error || "Error al iniciar sesión");
      }

      console.log(data);
      localStorage.setItem("token", data.token);
      localStorage.setItem("userId", data.user.ID); // Asumiendo que el backend devuelve el id del usuario en data.user.id

      if (typeof onLogin === "function") {
        onLogin(data.user);
      } else {
        // Si no hay función onLogin, redirigir directamente
        window.location.href = "/miscursos";
      }
    } catch (err) {
      if (err.message === "Failed to fetch") {
        setError(
          "No se pudo conectar al servidor. Por favor, verifique que el servidor esté corriendo.",
        );
      } else {
        setError(err.message);
      }
      console.error("Error durante el login:", err);
    }
  };

  return (
    <div className="login-container">
      <h2 className="login-title">Iniciar Sesión</h2>
      {error && <div className="login-error">{error}</div>}
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
      <button onClick={handleLogin} className="login-button">
        Ingresar
      </button>
    </div>
  );
};

export default Login;
