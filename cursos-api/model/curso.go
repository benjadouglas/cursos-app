package model

type Curso struct {
	Id        string  `bson:"_id,omitempty"`
	Nombre    string  `bson:"Nombre"`
	Precio    float64 `bson:"Precio"`
	Profesor  string  `bson:"Profesor"`
	Capacidad int     `bson:"Capacidad"`
	Duracion  string  `bson:"Duracion"`
}
