package solr

import (
	"context"
	"encoding/json"
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

func (this Solr) Search(ctx context.Context, query string, offset int, limit int) {
	queryParser := solr.NewDisMaxQueryParser().Query(query).BuildParser()
	solrQuery := solr.NewQuery(queryParser).
		Sort("rating").
		Offset(offset).
		Limit(limit)
	// fmt.Printf("Solr Query: %+v\n", solrQuery.BuildQuery())
	response, err := this.Client.Query(ctx, this.Collection, solrQuery)
	FailOnErr(err, "Error getting query\n")
	bytes, err := json.Marshal(response.Response.Documents)
	FailOnErr(err, "Error Marshalling response\n")
	result := make([]models.Curso, 0)
	if err := json.Unmarshal(bytes, &result); err != nil {
		FailOnErr(err, "Error Unmarshalling response\n")
	}
	println(result[0].Id)
}
