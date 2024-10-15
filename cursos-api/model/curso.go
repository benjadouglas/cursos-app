package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Curso struct {
	ID     primitive.ObjectID `bson:"_id"`
	Nombre string
	Precio float64
}
