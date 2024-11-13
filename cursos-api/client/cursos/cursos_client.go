package cursos

// "context"
// "cursos-api/db"
// "cursos-api/model"
// "cursos-api/rabbit"

// log "github.com/sirupsen/logrus"

// "go.mongodb.org/mongo-driver/bson"
// "go.mongodb.org/mongo-driver/bson/primitive"

// var client = db.Connect()

// func GetCursos() ([]model.Curso, error) {
// 	var cursos []model.Curso
// 	coll := client.Database("db").Collection("cursos")
// 	cur, err := coll.Find(context.TODO(), bson.M{})
// 	if err != nil {
// 		log.Error(err)
// 		return nil, err
// 	}
// 	defer cur.Close(context.TODO())
// 	for cur.Next(context.TODO()) {
// 		var curso model.Curso
// 		err := cur.Decode(&curso)
// 		if err != nil {
// 			log.Error(err)
// 			return nil, err
// 		}
// 		cursos = append(cursos, curso)
// 	}

// 	if err := cur.Err(); err != nil {
// 		log.Error(err)
// 		return nil, err
// 	}

// 	return cursos, nil
// }

// func GetCursoById(id string) (model.Curso, error) {
// 	var curso model.Curso
// 	_id, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return model.Curso{}, err
// 	}
// 	coll := client.Database("db").Collection("cursos")
// 	_err := coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: _id}}).Decode(&curso)
// 	if _err != nil {
// 		return model.Curso{}, _err
// 	}
// 	return curso, nil
// }

// func EditCurso(id string, curso model.Curso) error {
// 	_id, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return err
// 	}
// 	coll := client.Database("db").Collection("cursos")
// 	filter := bson.D{primitive.E{Key: "_id", Value: _id}}
// 	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
// 		{Key: "Nombre", Value: curso.Nombre},
// 		{Key: "Precio", Value: curso.Precio},
// 	}}}
// 	_, err = coll.UpdateOne(context.TODO(), filter, update)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	curso.ID = _id
// 	rabbit.Publish(curso)

// 	return nil
// }

// func DeleteCurso(id string) error {
// 	_id, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return err
// 	}

// 	curso, err := GetCursoById(id)
// 	if err != nil {
// 		return err
// 	}

// 	coll := client.Database("db").Collection("cursos")
// 	_, err = coll.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: _id}})
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	rabbit.Publish(curso)

// 	return nil
// }

// func CreateCurso(curso model.Curso) error {
// 	coll := client.Database("db").Collection("cursos")
// 	curso.ID = primitive.NewObjectID()
// 	_, err := coll.InsertOne(context.TODO(), curso)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	rabbit.Publish(curso)

// 	return nil
// }
