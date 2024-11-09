## Queries

Search for Id:

- curl http://localhost:8983/solr/cursos/query\?q\=Id:2

---

# QUERY NOT SENDING

## Error 1: undefined field name

```
› go run main.go

2024/11/09 15:01:18 QSY failed to execute search query: undefined field name
panic: QSY failed to execute search query: undefined field name
```

### Problem

- The query isnt right

  - What exacty isnt right?
  - Can we see the query being executed?

- Lets try to see the path of the query

```
q=(name:Id:2)&rows=0&start=102024/11/09
```

- This was the query that it was executing
- I dont know why there is a date, we delete that and we augment the number of rows that can return

- **Final query:**

```
q=Id:[id]$rows=[rows]
```

## Error 2: undefined field again

```
› go run main.go

2024/11/09 15:28:32 Error on failonerror N°2 undefined field q=Id
panic: Error on failonerror N°2 undefined field q=Id
```

- Undefined field q=Id

### Problem

- The parser is complaining?
- It may be that the initial part of the url isnt right, the whole http...

## Fix

- Rewrited the whole Search() method again
- Just set an Id:2 query
- The q= part of the query was causing the error so i deleted it
- There was a problem with the client

```
func (searchEngine Solr) Search() {
	query := fmt.Sprintf("Id:2")


	resp, err := searchEngine.Client.Query(context.Background(), "cursos", solr.NewQuery(query))
	FailOnErr(err, "The query failed sending")
	FailOnErr(resp.Error, "The query returned with an error")

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
```
