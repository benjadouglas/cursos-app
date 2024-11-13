package model

type Curso struct {
	Id        string  `json:"id"`
	Nombre    string  `json:"Nombre"`
	Precio    float64 `json:"Precio"`
	Profesor  string  `json:"Profesor"`
	Capacidad int     `json:"Capacidad"`
	Duracion  string  `json:"Duracion"`
}
