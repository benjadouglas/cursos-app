package cursos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	Dao "search-api/dao"

	"github.com/sirupsen/logrus"
	"github.com/stevenferrer/solr-go"
)

type SolrConfig struct {
	BaseURL    string
	Collection string
}

type Solr struct {
	Client     *solr.JSONClient
	Collection string
}

func NewSolr(config SolrConfig) Solr {
	return Solr{
		Client:     solr.NewJSONClient(config.BaseURL),
		Collection: config.Collection,
	}
}

func (searchEngine Solr) Search(ctx context.Context, query string, limit int, offset int) ([]Dao.Curso, error) {

	logrus.Printf(query)
	solrQuery := solr.NewQuery(query).Limit(limit).Offset(offset)
	resp, err := searchEngine.Client.Query(ctx, "courses", solrQuery)

	logrus.Printf("%v", resp)
	logrus.Printf("%v", err)
	if err != nil {
		return nil, err
	}

	var coursesList []Dao.Curso
	for _, doc := range resp.Response.Documents {
		course := Dao.Curso{
			Id:        getStringField(doc, "Id"),
			Nombre:    getStringField(doc, "Nombre"),
			Precio:    getFloatField(doc, "Precio"),
			Profesor:  getStringField(doc, "Profesor"),
			Capacidad: getIntField(doc, "Capacidad"),
			Duracion:  getStringField(doc, "Duracion"),
		}
		coursesList = append(coursesList, course)
	}
	return coursesList, nil

}

func (searchEngine Solr) Update(ctx context.Context, curso Dao.Curso) error {
	doc := map[string]interface{}{
		"Id":        curso.Id,
		"Nombre":    curso.Nombre,
		"Precio":    curso.Precio,
		"Profesor":  curso.Profesor,
		"Capacidad": curso.Capacidad,
		"Duracion":  curso.Duracion,
	}

	updateRequest := map[string]interface{}{
		"add": []interface{}{doc},
	}

	body, err := json.Marshal(updateRequest)
	if err != nil {
		return fmt.Errorf("error marshaling curso document: %w", err)
	}

	resp, err := searchEngine.Client.Update(ctx, searchEngine.Collection, solr.JSON, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("error updating curso: %w", err)
	}
	if resp.Error != nil {
		return fmt.Errorf("failed to update curso: %v", resp.Error)
	}

	if err := searchEngine.Client.Commit(ctx, searchEngine.Collection); err != nil {
		return fmt.Errorf("error committing changes to Solr: %w", err)
	}

	return nil
}

func (searchEngine Solr) Delete(ctx context.Context, id string) error {
	docToDelete := map[string]interface{}{
		"delete": map[string]interface{}{
			"id": id,
		},
	}

	body, err := json.Marshal(docToDelete)
	if err != nil {
		return fmt.Errorf("error marshaling curso document: %w", err)
	}

	resp, err := searchEngine.Client.Update(ctx, searchEngine.Collection, solr.JSON, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("error deleting curso: %w", err)
	}
	if resp.Error != nil {
		return fmt.Errorf("failed to delete curso: %v", resp.Error)
	}
	if err := searchEngine.Client.Commit(ctx, searchEngine.Collection); err != nil {
		return fmt.Errorf("error committing changes to Solr: %w", err)
	}

	return nil
}

func (searchEngine Solr) Index(ctx context.Context, curso Dao.Curso) (string, error) {
	doc := map[string]interface{}{
		"Id":        curso.Id,
		"Nombre":    curso.Nombre,
		"Precio":    curso.Precio,
		"Profesor":  curso.Profesor,
		"Capacidad": curso.Capacidad,
		"Duracion":  curso.Duracion,
	}

	indexRequest := map[string]interface{}{
		"add": []interface{}{doc},
	}

	body, err := json.Marshal(indexRequest)
	if err != nil {
		return "", fmt.Errorf("error marshaling curso document: %w", err)
	}

	resp, err := searchEngine.Client.Update(ctx, searchEngine.Collection, solr.JSON, bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("error indexing curso: %w", err)
	}
	if resp.Error != nil {
		return "", fmt.Errorf("failed to index curso: %v", resp.Error)
	}

	if err := searchEngine.Client.Commit(ctx, searchEngine.Collection); err != nil {
		return "", fmt.Errorf("error committing changes to Solr: %w", err)
	}

	return curso.Id, nil
}

func getStringField(doc map[string]interface{}, field string) string {
	if val, ok := doc[field].(string); ok {
		return val
	}
	if val, ok := doc[field].([]interface{}); ok && len(val) > 0 {
		if strVal, ok := val[0].(string); ok {
			return strVal
		}
	}
	return ""
}

func getFloatField(doc map[string]interface{}, field string) float64 {
	if val, ok := doc[field].(float64); ok {
		return val
	}
	if val, ok := doc[field].([]interface{}); ok && len(val) > 0 {
		if floatVal, ok := val[0].(float64); ok {
			return floatVal
		}
	}
	return 0.0
}

func getIntField(doc map[string]interface{}, field string) int {
	if val, ok := doc[field].(float64); ok {
		return int(val)
	}
	if val, ok := doc[field].([]interface{}); ok && len(val) > 0 {
		if floatVal, ok := val[0].(float64); ok {
			return int(floatVal)
		}
	}
	return 0
}
