package models

type Curso struct {
	Id     string    `json:"Id"`
	Nombre []string  `json:"Nombre"`
	Precio []float64 `json:"Precio"`
}
