package cursos

import (
	"context"
	domain "cursos-api/domain/cursos"
	"cursos-api/model"
	"fmt"
)

type Repository interface {
	GetCursoByID(ctx context.Context, id string) (model.Curso, error)
	Update(ctx context.Context, curso model.Curso) error
	Create(ctx context.Context, curso model.Curso) (string, error)
	Delete(ctx context.Context, id string) error
}

type Queue interface {
	Publish(curso domain.CursoNew) error
}

type Service struct {
	mainRepository Repository
	eventsQueue    Queue
}

func NewService(mainRepository Repository, eventsQueue Queue) Service {
	return Service{
		mainRepository: mainRepository,
		eventsQueue:    eventsQueue,
	}
}

func (service Service) GetCursoByID(ctx context.Context, id string) (domain.Curso, error) {
	curso, err := service.mainRepository.GetCursoByID(ctx, id)
	if err != nil {
		return domain.Curso{}, fmt.Errorf("error getting curso from repository: %v", err)
	}
	return domain.Curso{
		Id:        curso.Id,
		Nombre:    curso.Nombre,
		Precio:    curso.Precio,
		Profesor:  curso.Profesor,
		Capacidad: curso.Capacidad,
		Duracion:  curso.Duracion,
		Maximo:    curso.Maximo,
	}, nil
}

func (service Service) Create(ctx context.Context, curso domain.Curso) (string, error) {
	record := model.Curso{
		Id:        curso.Id,
		Nombre:    curso.Nombre,
		Precio:    curso.Precio,
		Profesor:  curso.Profesor,
		Capacidad: curso.Capacidad,
		Duracion:  curso.Duracion,
		Maximo:    curso.Maximo,
	}

	Id, err := service.mainRepository.Create(ctx, record)
	if err != nil {
		return "", fmt.Errorf("error creating curso in main repository: %w", err)
	}
	record.Id = Id
	if err := service.eventsQueue.Publish(domain.CursoNew{
		Operation: "CREATE",
		CursoID:   Id,
	}); err != nil {
		return "", fmt.Errorf("error publishing curso new: %w", err)
	}

	return Id, nil
}

func (service Service) Update(ctx context.Context, curso domain.Curso) error {
	record := model.Curso{
		Id:        curso.Id,
		Nombre:    curso.Nombre,
		Precio:    curso.Precio,
		Profesor:  curso.Profesor,
		Capacidad: curso.Capacidad,
		Duracion:  curso.Duracion,
		Maximo:    curso.Maximo,
	}
	err := service.mainRepository.Update(ctx, record)
	if err != nil {
		return fmt.Errorf("error updating curso in repository: %w", err)
	}
	if err := service.eventsQueue.Publish(domain.CursoNew{
		Operation: "UPDATE",
		CursoID:   curso.Id,
	}); err != nil {
		return fmt.Errorf("error publishing curso update: %w", err)
	}
	return nil
}

func (service Service) Delete(ctx context.Context, id string) error {
	// Delete from repository
	err := service.mainRepository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("error deleting curso from repository: %w", err)
	}

	// Publish delete event
	if err := service.eventsQueue.Publish(domain.CursoNew{
		Operation: "DELETE",
		CursoID:   id,
	}); err != nil {
		return fmt.Errorf("error publishing curso delete: %w", err)
	}

	return nil
}
