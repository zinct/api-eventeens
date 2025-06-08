package usecases

import (
	"context"
	"errors"
	"goevents/internal/domain/entities"
	"goevents/internal/domain/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestEventUsecase_Create_Success(t *testing.T) {
	repoMock := new(repositories.EventRepositoryMock)
	usecase := NewEventUsecase(repoMock)

	repoMock.On("Create", mock.Anything, mock.Anything).Return(nil)

	err := usecase.Create(context.TODO(), &entities.Event{
		ID:          "1",
		Title:       "Test Event",
		Description: "Test Description",
	})

	assert.NoError(t, err)
}

func TestEventUsecase_Create_Error(t *testing.T) {
	repoMock := new(repositories.EventRepositoryMock)
	usecase := NewEventUsecase(repoMock)

	repoMock.On("Create", mock.Anything, mock.Anything).Return(errors.New("error"))

	err := usecase.Create(context.TODO(), &entities.Event{
		ID:          "1",
		Title:       "Test Event",
		Description: "Test Description",
	})

	assert.Error(t, err)
}

func TestEventUsecase_FindAll_Success(t *testing.T) {
	repoMock := new(repositories.EventRepositoryMock)
	usecase := NewEventUsecase(repoMock)

	events := []*entities.Event{
		{
			ID:          "1",
			Title:       "Event 1",
			Description: "Desc 1",
		},
		{
			ID:          "2",
			Title:       "Event 2",
			Description: "Desc 2",
		},
	}

	repoMock.On("FindAll", mock.Anything).Return(events, nil)

	result, err := usecase.FindAll(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, events, result)
}

func TestEventUsecase_FindAll_Error(t *testing.T) {
	repoMock := new(repositories.EventRepositoryMock)
	usecase := NewEventUsecase(repoMock)

	repoMock.On("FindAll", mock.Anything).Return(nil, errors.New("error"))

	result, err := usecase.FindAll(context.TODO())
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestEventUsecase_Find_Success(t *testing.T) {
	repoMock := new(repositories.EventRepositoryMock)
	usecase := NewEventUsecase(repoMock)
	event := &entities.Event{
		ID:          "1",
		Title:       "Event 1",
		Description: "Desc 1",
	}

	repoMock.On("Find", mock.Anything, "1").Return(event, nil)

	result, err := usecase.Find(context.TODO(), "1")
	assert.NoError(t, err)
	assert.Equal(t, event, result)
}

func TestEventUsecase_Find_Error(t *testing.T) {
	repoMock := new(repositories.EventRepositoryMock)
	usecase := NewEventUsecase(repoMock)

	repoMock.On("Find", mock.Anything, "1").Return(nil, errors.New("error"))

	result, err := usecase.Find(context.TODO(), "1")
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestEventUsecase_Update_Success(t *testing.T) {
	repoMock := new(repositories.EventRepositoryMock)
	usecase := NewEventUsecase(repoMock)
	event := &entities.Event{
		ID:          "1",
		Title:       "Event 1",
		Description: "Desc 1",
	}

	repoMock.On("Update", mock.Anything, event).Return(nil)

	err := usecase.Update(context.TODO(), event)
	assert.NoError(t, err)
}

func TestEventUsecase_Update_Error(t *testing.T) {
	repoMock := new(repositories.EventRepositoryMock)
	usecase := NewEventUsecase(repoMock)
	event := &entities.Event{
		ID:          "1",
		Title:       "Event 1",
		Description: "Desc 1",
	}

	repoMock.On("Update", mock.Anything, event).Return(errors.New("error"))

	err := usecase.Update(context.TODO(), event)
	assert.Error(t, err)
}

func TestEventUsecase_Delete_Success(t *testing.T) {
	repoMock := new(repositories.EventRepositoryMock)
	usecase := NewEventUsecase(repoMock)

	repoMock.On("Delete", mock.Anything, "1").Return(nil)

	err := usecase.Delete(context.TODO(), "1")
	assert.NoError(t, err)
}

func TestEventUsecase_Delete_Error(t *testing.T) {
	repoMock := new(repositories.EventRepositoryMock)
	usecase := NewEventUsecase(repoMock)

	repoMock.On("Delete", mock.Anything, "1").Return(errors.New("error"))

	err := usecase.Delete(context.TODO(), "1")
	assert.Error(t, err)
}
