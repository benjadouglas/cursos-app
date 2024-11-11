package cursos

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	cursosModel "search-api/models"
)

type HTTPConfig struct {
	Host string
	Port string
}

type HTTP struct {
	baseURL func(cursoId string) string
}

func NewHTTP(config HTTPConfig) HTTP {
	return HTTP{
		baseURL: func(cursoId string) string {
			return fmt.Sprintf("http://%s:%s/cursos/%s", config.Host, config.Port, cursoId)
		},
	}
}

func (repository HTTP) GetCursoByID(ctx context.Context, id string) (cursosModel.Curso, error) {
	resp, err := http.Get(repository.baseURL(id))
	if err != nil {
		return cursosModel.Curso{}, fmt.Errorf("Error fetching curso (%s): %w\n", id, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return cursosModel.Curso{}, fmt.Errorf("Failed to fetch curso (%s): received status code %d\n", id, resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return cursosModel.Curso{}, fmt.Errorf("Error reading response body for curso (%s): %w\n", id, err)
	}

	// Unmarshal the curso details into the curso struct
	var curso cursosModel.Curso
	if err := json.Unmarshal(body, &curso); err != nil {
		return cursosModel.Curso{}, fmt.Errorf("Error unmarshaling curso data (%s): %w\n", id, err)
	}

	return curso, nil
}
