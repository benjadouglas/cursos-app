import React, { useState } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './components/Home';
import Resultados from './components/Resultados';
import Detalle from './components/Detalle';
import Login from './components/Login';
import Registro from './components/Registro';
import Congrats from './components/Congrats';
import MisCursos from './components/MisCursos';
import Cursos from './components/Cursos';

const App = () => {
    const [cursos, setCursos] = useState([]);

    return (
        <Router>
            <Routes>
                <Route path="/" element={<Home setCursos={setCursos} />} />
                <Route path="/resultados" element={<Resultados cursos={cursos} />} />
                <Route path="/detalle" element={<Detalle />} />
                <Route path="/login" element={<Login />} />
                <Route path="/registro" element={<Registro />} />
                <Route path="/congrats" element={<Congrats />} />
                <Route path="/miscursos" element={<MisCursos />} />
                <Route path="/cursos" element={<Cursos />} />
            </Routes>
        </Router>
    );
};

export default App;