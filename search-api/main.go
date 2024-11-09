package main

import (
	"search-api/solr"
)

func main() {
	config := solr.SolrConfig{
		BaseURL:    "http://localhost:8983",
		Collection: "cursos",
	}
	solaris := solr.NewSolr(config)
	solaris.Search()
}
