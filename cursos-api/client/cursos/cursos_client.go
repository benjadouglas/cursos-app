package cursos

import (
	"context"
	"cursos-api/db"
	"cursos-api/model"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client = db.Connect()

func GetCursos() ([]model.Curso, error) {
	var cursos []model.Curso
	coll := client.Database("db").Collection("cursos")
	cur, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var curso model.Curso
		err := cur.Decode(&curso)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		cursos = append(cursos, curso)
	}

	if err := cur.Err(); err != nil {
		log.Error(err)
		return nil, err
	}

	return cursos, nil
}

func GetCursoById(id string) (model.Curso, error) {
	var curso model.Curso
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Curso{}, err
	}
	coll := client.Database("db").Collection("cursos")
	_err := coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: _id}}).Decode(&curso)
	if _err != nil {
		return model.Curso{}, _err

	}
	return curso, nil
}
