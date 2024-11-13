import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./Registro.css";

const Registro = () => {
  const [formData, setFormData] = useState({
    nombre: "",
    email: "",
    password: "",
  });
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
    setLoading(true);

    // Validaciones básicas
    if (!formData.nombre || !formData.email || !formData.password) {
      setError("Todos los campos son obligatorios");
      setLoading(false);
      return;
    }

    try {
      const response = await fetch("http://localhost:8085/api/users", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.error || "Error en el registro");
      }

      // Registro exitoso
      navigate("/login");
    } catch (err) {
      setError(err.message);
      console.error("Error durante el registro:", err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="registro-container">
      <h2 className="registro-title">Registro de Usuario</h2>
      {error && <div className="registro-error">{error}</div>}
      <form onSubmit={handleSubmit} className="registro-form">
        <input
          type="text"
          name="nombre"
          value={formData.nombre}
          onChange={(e) => setFormData({ ...formData, nombre: e.target.value })}
          className="registro-input"
          placeholder="Nombre"
        />
        <input
          type="email"
          name="email"
          value={formData.email}
          onChange={(e) => setFormData({ ...formData, email: e.target.value })}
          className="registro-input"
          placeholder="Email"
        />
        <input
          type="password"
          name="password"
          value={formData.password}
          onChange={(e) =>
            setFormData({ ...formData, password: e.target.value })
          }
          className="registro-input"
          placeholder="Contraseña"
        />
        <button type="submit" className="registro-button" disabled={loading}>
          {loading ? "Registrando..." : "Registrarse"}
        </button>
      </form>
    </div>
  );
};

export default Registro;
