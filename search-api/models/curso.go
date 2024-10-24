package models

type Curso struct {
	Id     int8   `bson:"Id,omitempty"`
	Nombre string `bson:"Nombre,omitempty"`
	Precio int16  `bson:"Precio,omitempty"`
}
