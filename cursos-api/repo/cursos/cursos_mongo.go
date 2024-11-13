package cursos

import (
	"context"
	"cursos-api/model"
	"fmt"
	"log"
	"strings"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type MongoConfig struct {
	Host       string
	Port       string
	Username   string
	Password   string
	Database   string
	Collection string
}

type Mongo struct {
	client     *mongo.Client
	database   string
	collection string
}

const (
	connectionURI = "mongodb://%s:%s"
)

func NewMongo(config MongoConfig) Mongo {
	credentials := options.Credential{
		Username: config.Username,
		Password: config.Password,
	}

	ctx := context.Background()
	uri := fmt.Sprintf(connectionURI, config.Host, config.Port)
	cfg := options.Client().ApplyURI(uri).SetAuth(credentials)

	client, err := mongo.Connect(ctx, cfg)
	if err != nil {
		log.Panicf("error connecting to mongo DB: %v", err)
	}

	return Mongo{
		client:     client,
		database:   config.Database,
		collection: config.Collection,
	}
}

func (repository Mongo) GetCursoByID(ctx context.Context, id string) (model.Curso, error) {
	// Get from MongoDB
	logrus.Printf("HERHERR: %s", id)
	// Trim the "id:" prefix if present
	id = strings.TrimPrefix(id, "id:")
	logrus.Printf("Trimmed ID: %s", id)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Curso{}, fmt.Errorf("error converting id to mongo ID: %w", err)
	}
	result := repository.client.Database(repository.database).Collection(repository.collection).FindOne(ctx, bson.M{"_id": objectID})
	if result.Err() != nil {
		return model.Curso{}, fmt.Errorf("error finding document: %w", result.Err())
	}

	// Convert document to DAO
	var cursoDao model.Curso
	if err := result.Decode(&cursoDao); err != nil {
		return model.Curso{}, fmt.Errorf("error decoding result: %w", err)
	}
	return cursoDao, nil
}

func (repository Mongo) Create(ctx context.Context, curso model.Curso) (string, error) {
	// Insert into mongo
	result, err := repository.client.Database(repository.database).Collection(repository.collection).InsertOne(ctx, curso)
	if err != nil {
		return "", fmt.Errorf("error creating document: %w", err)
	}

	// Get inserted ID
	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("error converting mongo ID to object ID")
	}
	return objectID.Hex(), nil
}

func (repository Mongo) Update(ctx context.Context, curso model.Curso) error {
	// Convert curso ID to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(curso.ID)
	if err != nil {
		return fmt.Errorf("error converting id to mongo ID: %w", err)
	}

	update := bson.M{}

	if curso.Nombre != "" {
		update["nombre"] = curso.Nombre
	}
	if curso.Precio != 0 {
		update["precio"] = curso.Precio
	}
	if curso.Profesor != "" {
		update["profesor"] = curso.Profesor
	}
	if curso.Capacidad != 0 {
		update["capacidad"] = curso.Capacidad
	}
	if curso.Duracion != "" {
		update["duracion"] = curso.Duracion
	}

	if len(update) == 0 {
		return fmt.Errorf("no fields to update for curso ID %s", curso.ID)
	}

	filter := bson.M{"_id": objectID}
	result, err := repository.client.Database(repository.database).Collection(repository.collection).UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return fmt.Errorf("error updating document: %w", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("no document found with ID %s", curso.ID)
	}

	return nil
}

func (repository Mongo) Delete(ctx context.Context, id string) error {
	// Convert curso ID to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("error converting id to mongo ID: %w", err)
	}
	// Delete the document from MongoDB
	filter := bson.M{"_id": objectID}
	result, err := repository.client.Database(repository.database).Collection(repository.collection).DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("error deleting document: %w", err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("no document found with ID %s", id)
	}

	return nil
}
