package search

import (
	"context"
	"fmt"
	cursoDAO "search-api/dao"
	cursoModel "search-api/models"
)

type Repository interface {
	Index(ctx context.Context, curso cursoDAO.Curso) (string, error)
	Update(ctx context.Context, curso cursoDAO.Curso) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, query string, limit int, offset int) ([]cursoDAO.Curso, error)
}

type ExternalRepository interface {
	GetCursoByID(ctx context.Context, id string) (cursoModel.Curso, error)
}

type Queue interface {
	Publish(cursoNew cursoModel.CursoNew) error
}

type Service struct {
	repository  Repository
	cursosAPI   ExternalRepository
	eventsQueue Queue
}

func NewService(repository Repository, cursosAPI ExternalRepository, eventsQueue Queue) Service {
	return Service{
		repository:  repository,
		cursosAPI:   cursosAPI,
		eventsQueue: eventsQueue,
	}
}

func (service Service) Search(ctx context.Context, query string, offset int, limit int) ([]cursoModel.Curso, error) {
	cursoDaoLIst, err := service.repository.Search(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error searching cursos: %w", err)
	}

	cursoModelList := make([]cursoModel.Curso, 0)
	for _, curso := range cursoDaoLIst {
		cursoModelList = append(cursoModelList, cursoModel.Curso{
			Id:        curso.Id,
			Nombre:    curso.Nombre,
			Precio:    curso.Precio,
			Profesor:  curso.Profesor,
			Capacidad: curso.Capacidad,
			Duracion:  curso.Duracion,
			Maximo:    curso.Maximo,
		})
	}

	return cursoModelList, nil
}

func (service Service) HandleCursoNew(cursoNew cursoModel.CursoNew) {
	switch cursoNew.Operation {
	case "CREATE", "UPDATE":
		curso, err := service.cursosAPI.GetCursoByID(context.Background(), cursoNew.CursoId)
		if err != nil {
			fmt.Printf("Error getting curso (%s) from API: %v\n", cursoNew.CursoId, err)
			return
		}

		cursoDAO := cursoDAO.Curso{
			Id:        curso.Id,
			Nombre:    curso.Nombre,
			Precio:    curso.Precio,
			Profesor:  curso.Profesor,
			Capacidad: curso.Capacidad,
			Duracion:  curso.Duracion,
			Maximo:    curso.Maximo,
		}
		fmt.Printf("cursoDAO: %v \n", cursoDAO)
		if cursoNew.Operation == "CREATE" {
			if _, err := service.repository.Index(context.Background(), cursoDAO); err != nil {
				fmt.Printf("Error indexing curso (%s): %v\n", cursoNew.CursoId, err)
			} else {
				fmt.Println("Curso indexed successfully:", cursoNew.CursoId)
			}
		} else {
			if err := service.repository.Update(context.Background(), cursoDAO); err != nil {
				fmt.Printf("Error updating curso (%s): %v\n", cursoNew.CursoId, err)
			} else {
				fmt.Println("Curso updated successfully:", cursoNew.CursoId)
			}
		}

	case "DELETE":
		if err := service.repository.Delete(context.Background(), cursoNew.CursoId); err != nil {
			fmt.Printf("Error deleting curso (%s): %v\n", cursoNew.CursoId, err)
		} else {
			fmt.Println("Curso deleted successfully:", cursoNew.CursoId)
		}

	default:
		fmt.Printf("Unknown operation: %s\n", cursoNew.Operation)
	}
}
