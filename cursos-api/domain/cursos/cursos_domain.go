package cursos

type Curso struct {
	Id        string  `json:"Id"`
	Nombre    string  `json:"Nombre"`
	Precio    float64 `json:"Precio"`
	Profesor  string  `json:"Profesor"`
	Capacidad int     `json:"Capacidad"`
	Duracion  string  `json:"Duracion"`
	Maximo    int     `json:"Maximo"`
}

type CursoNew struct {
	Operation string `json:"operation"`
	CursoID   string `json:"curso_id"`
}
