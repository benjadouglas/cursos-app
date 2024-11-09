package solr

import (
	"context"
	"fmt"
	"search-api/models"
	. "search-api/utils"

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

func (searchEngine Solr) Search() {
	baseURL := "http://localhost:8983"
	client := solr.NewJSONClient(baseURL)

	query := solr.NewQuery(fmt.Sprintf("Id:2"))

	resp, err := client.Query(context.TODO(), "cursos", query)

	FailOnErr(err, "The query failed sending")

	var coursesList []models.Curso
	for _, doc := range resp.Response.Documents {
		course := models.Curso{
			Id:     getStringField(doc, "Id"),
			Nombre: []string{getStringField(doc, "Nombre")},
			Precio: []float64{getFloatField(doc, "Precio")},
		}
		coursesList = append(coursesList, course)
	}
	fmt.Printf("Courses found: %+v\n", coursesList)

}

// func (searchEngine Solr) Search(ctx context.Context, query string, limit int) ([]models.Curso, error) {
//
// 	// Prepare the Solr query with limit and offset
// 	solrQuery := solr.NewQuery(query string)

// 	// Execute the search request
// 	resp, err := searchEngine.Client.Query(ctx, searchEngine.Collection, solr.NewQuery(solrQuery))

// 	// N°1
// 	FailOnErr(err, "Error on failonerror N°1")

// 	// N°2
// 	FailOnErr(resp.Error, "Error on failonerror N°2")

// 	// Parse the response and extract course documents
// 	var coursesList []models.Curso
// 	for _, doc := range resp.Response.Documents {
// 		// Safely extract course fields with type assertions
// 		course := models.Curso{
// 			Id:     getStringField(doc, "Id"),
// 			Nombre: []string{getStringField(doc, "Nombre")},
// 			Precio: []float64{getFloatField(doc, "Precio")},
// 		}
// 		coursesList = append(coursesList, course)
// 	}

// 	return coursesList, nil
// }

// func (this Solr) Search(ctx context.Context, query string, offset int, limit int) {
// 	queryParser := solr.NewDisMaxQueryParser().Query(query).BuildParser()
// 	solrQuery := solr.NewQuery(queryParser).
// 		Sort("rating").
// 		Offset(offset).
// 		Limit(limit)

// 	response, err := this.Client.Query(ctx, this.Collection, solrQuery)
// 	FailOnErr(err, "Error getting query\n")

// 	bytes, err := json.Marshal(response)
// 	FailOnErr(err, "Error Marshalling response\n")

// 	var solrResponse SolrResponse
// 	err1 := json.Unmarshal(bytes, &solrResponse)
// 	FailOnErr(err1, "Error unmarshaling response")

// 	result := make([]models.Curso, 0)
// 	if err := json.Unmarshal(bytes, &result); err != nil {
// 		FailOnErr(err, "Error Unmarshalling response\n")
// 	}
// 	println(result[0].Id)

// }

// Helper function to safely get string fields from the document
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

// Helper function to safely get float64 fields from the document
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
