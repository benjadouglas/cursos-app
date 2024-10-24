package main

import (
	"context"
	"search-api/solr"
)

func main() {
	config := solr.SolrConfig{
		BaseURL:    "http://localhost:8983",
		Collection: "cursos",
	}
	solaris := solr.NewSolr(config)
	solaris.Search(context.TODO(), "q=Id:2", 10, 10)
}
