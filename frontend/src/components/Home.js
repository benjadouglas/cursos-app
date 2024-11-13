import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import './Home.css';

const Home = ({ setCursos }) => {
    const [search, setSearch] = useState('');
    const navigate = useNavigate();

    const cursosDisponibles = [
        { id: 1, nombre: "Curso de Matemáticas", descripcion: "Aprende matemáticas desde cero", puntuacion: 4.5 },
        { id: 2, nombre: "Curso de Historia", descripcion: "Historia universal", puntuacion: 4.2 },
        { id: 3, nombre: "Curso de Ciencias", descripcion: "Explora el mundo de la ciencia", puntuacion: 4.7 },
    ];

    const handleSearch = () => {
        const cursosFiltrados = cursosDisponibles.filter(curso =>
            curso.nombre.toLowerCase().includes(search.toLowerCase())
        );
        setCursos(cursosFiltrados);
        navigate('/resultados');
    };

    return (
        <div className="home-container">
            <h1 className="home-title">Bienvenido a la Plataforma de Cursos</h1>
            <p className="home-subtitle">Encuentra y aprende a tu ritmo</p>

            <div className="home-search">
                <input
                    type="text"
                    placeholder="Buscar cursos..."
                    value={search}
                    onChange={(e) => setSearch(e.target.value)}
                    className="search-input"
                />
                <button onClick={handleSearch} className="search-button">Buscar</button>
            </div>

            <div className="home-buttons">
                <Link to="/registro" className="home-link">
                    <button className="action-button">Registrarse</button>
                </Link>
                <Link to="/login" className="home-link">
                    <button className="action-button">Iniciar Sesión</button>
                </Link>
                <Link to="/cursos" className="home-link">
                    <button className="action-button">Ver Cursos</button>
                </Link>
            </div>

            <div className="home-bottom-image"></div>
        </div>
    );
};

export default Home;