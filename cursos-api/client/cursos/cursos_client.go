package cursos

import (
	"context"
	"cursos-api/db"
	"cursos-api/model"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
)

var client = db.Connect()

func GetAllCursos() []model.Curso {
	var cursos []model.Curso
	coll := client.Database("db").Collection("cursos")
	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return cursos
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var curso model.Curso
		err := cursor.Decode(&curso)
		if err != nil {
			continue
		}
		cursos = append(cursos, curso)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursos retrieved: %+v", cursos)
		return cursos
	}

	log.Printf("Cursos retrieved: %+v", cursos)

	return cursos
}
