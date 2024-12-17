package cursos

import (
	"context"
	"cursos-api/domain/cursos"
	"cursos-api/model"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository simula el comportamiento de la interfaz Repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetCursoByID(ctx context.Context, id string) (model.Curso, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(model.Curso), args.Error(1)
}

func (m *MockRepository) Create(ctx context.Context, curso model.Curso) (string, error) {
	args := m.Called(ctx, curso)
	return args.String(0), args.Error(1)
}

func (m *MockRepository) Update(ctx context.Context, curso model.Curso) error {
	args := m.Called(ctx, curso)
	return args.Error(0)
}

func (m *MockRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// MockQueue simula el comportamiento de la interfaz Queue
type MockQueue struct {
	mock.Mock
}

func (m *MockQueue) Publish(curso cursos.CursoNew) error {
	args := m.Called(curso)
	return args.Error(0)
}

// Test para GetCursoByID
func TestService_GetCursoByID(t *testing.T) {
	mockRepo := new(MockRepository)
	mockQueue := new(MockQueue)
	service := NewService(mockRepo, mockQueue)

	expectedCurso := model.Curso{
		Id:        "123",
		Nombre:    "Go Programming",
		Precio:    100,
		Profesor:  "John Doe",
		Capacidad: 30,
		Duracion:  "10h",
	}

	mockRepo.On("GetCursoByID", mock.Anything, "123").Return(expectedCurso, nil)

	ctx := context.Background()
	result, err := service.GetCursoByID(ctx, "123")

	assert.NoError(t, err)
	assert.Equal(t, "123", result.Id)
	assert.Equal(t, "Go Programming", result.Nombre)
	mockRepo.AssertExpectations(t)
}

// Test para Create
func TestService_Create(t *testing.T) {
	mockRepo := new(MockRepository)
	mockQueue := new(MockQueue)
	service := NewService(mockRepo, mockQueue)

	newCurso := cursos.Curso{
		Id:        "123",
		Nombre:    "Go Programming",
		Precio:    100,
		Profesor:  "John Doe",
		Capacidad: 30,
		Duracion:  "10h",
	}

	modelCurso := model.Curso{
		Id:        "123",
		Nombre:    "Go Programming",
		Precio:    100,
		Profesor:  "John Doe",
		Capacidad: 30,
		Duracion:  "10h",
	}

	mockRepo.On("Create", mock.Anything, modelCurso).Return("123", nil)
	mockQueue.On("Publish", cursos.CursoNew{
		Operation: "CREATE",
		CursoID:   "123",
	}).Return(nil)

	ctx := context.Background()
	id, err := service.Create(ctx, newCurso)

	assert.NoError(t, err)
	assert.Equal(t, "123", id)
	mockRepo.AssertExpectations(t)
	mockQueue.AssertExpectations(t)
}

// Test para Update
func TestService_Update(t *testing.T) {
	mockRepo := new(MockRepository)
	mockQueue := new(MockQueue)
	service := NewService(mockRepo, mockQueue)

	updateCurso := cursos.Curso{
		Id:        "123",
		Nombre:    "Advanced Go",
		Precio:    200,
		Profesor:  "Jane Doe",
		Capacidad: 25,
		Duracion:  "15h",
	}

	modelCurso := model.Curso{
		Id:        "123",
		Nombre:    "Advanced Go",
		Precio:    200,
		Profesor:  "Jane Doe",
		Capacidad: 25,
		Duracion:  "15h",
	}

	mockRepo.On("Update", mock.Anything, modelCurso).Return(nil)
	mockQueue.On("Publish", cursos.CursoNew{
		Operation: "UPDATE",
		CursoID:   "123",
	}).Return(nil)

	ctx := context.Background()
	err := service.Update(ctx, updateCurso)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	mockQueue.AssertExpectations(t)
}

// Test para Delete
func TestService_Delete(t *testing.T) {
	mockRepo := new(MockRepository)
	mockQueue := new(MockQueue)
	service := NewService(mockRepo, mockQueue)

	mockRepo.On("Delete", mock.Anything, "123").Return(nil)
	mockQueue.On("Publish", cursos.CursoNew{
		Operation: "DELETE",
		CursoID:   "123",
	}).Return(nil)

	ctx := context.Background()
	err := service.Delete(ctx, "123")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	mockQueue.AssertExpectations(t)
}

// Test para errores en GetCursoByID
func TestService_GetCursoByID_Error(t *testing.T) {
	mockRepo := new(MockRepository)
	mockQueue := new(MockQueue)
	service := NewService(mockRepo, mockQueue)

	mockRepo.On("GetCursoByID", mock.Anything, "123").Return(model.Curso{}, errors.New("repository error"))

	ctx := context.Background()
	_, err := service.GetCursoByID(ctx, "123")

	assert.Error(t, err)
	assert.EqualError(t, err, "error getting curso from repository: repository error")
	mockRepo.AssertExpectations(t)
}
