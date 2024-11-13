
import React from 'react';
import { Link } from 'react-router-dom';


const Navbar = () => {
    return (
        <nav className="navbar">
            <ul className="navbar-list">
                <li className="navbar-item">
                    <Link to="/">Home</Link>
                </li>
                <li className="navbar-item">
                    <Link to="/resultados">Resultados</Link>
                </li>
                <li className="navbar-item">
                    <Link to="/detalle">Detalle</Link>
                </li>
                <li className="navbar-item">
                    <Link to="/login">Login</Link>
                </li>
                <li className="navbar-item">
                    <Link to="/registro">Registro</Link>
                </li>
                <li className="navbar-item">
                    <Link to="/congrats">Congrats</Link>
                </li>
                <li className="navbar-item">
                    <Link to="/miscursos">Mis Cursos</Link>
                </li>
            </ul>
        </nav>
    );
};

export default Navbar;
