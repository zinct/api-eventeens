package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"goevents/internal/domain/entities"
	"goevents/internal/domain/usecases"
	"goevents/internal/interface/delivery/http/v1/requests"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupTest(t *testing.T) (*gin.Engine, *usecases.EventUsecaseMock, *EventController) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	usecaseMock := usecases.NewEventUsecaseMock()
	controller := NewEventController(usecaseMock)
	return router, usecaseMock, controller
}

func TestEventController_Create(t *testing.T) {
	router, repoMock, controller := setupTest(t)

	req := requests.CreateEventRequest{
		Title:       "Test Event",
		Description: "Test Description",
		Date:        time.Now(),
	}

	repoMock.On("Create", mock.Anything, mock.Anything).Return(nil)

	router.POST("/events", controller.Create)

	body, _ := json.Marshal(req)
	request, _ := http.NewRequest("POST", "/events", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	repoMock.AssertExpectations(t)
}

func TestEventController_Create_Error(t *testing.T) {
	router, repoMock, controller := setupTest(t)

	req := requests.CreateEventRequest{
		Title:       "Test Event",
		Description: "Test Description",
		Date:        time.Now(),
	}

	repoMock.On("Create", mock.Anything, mock.Anything).Return(errors.New("error"))

	router.POST("/events", controller.Create)

	body, _ := json.Marshal(req)
	request, _ := http.NewRequest("POST", "/events", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	repoMock.AssertExpectations(t)
}

func TestEventController_FindAll(t *testing.T) {
	router, repoMock, controller := setupTest(t)

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

	router.GET("/events", controller.FindAll)

	request, _ := http.NewRequest("GET", "/events", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	repoMock.AssertExpectations(t)
}

func TestEventController_FindAll_Error(t *testing.T) {
	router, repoMock, controller := setupTest(t)

	repoMock.On("FindAll", mock.Anything).Return(nil, errors.New("error"))

	router.GET("/events", controller.FindAll)

	request, _ := http.NewRequest("GET", "/events", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	repoMock.AssertExpectations(t)
}

func TestEventController_Find(t *testing.T) {
	router, repoMock, controller := setupTest(t)

	event := &entities.Event{
		ID:          "1",
		Title:       "Event 1",
		Description: "Desc 1",
	}

	repoMock.On("Find", mock.Anything, "1").Return(event, nil)

	router.GET("/events/:id", controller.Find)

	request, _ := http.NewRequest("GET", "/events/1", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	repoMock.AssertExpectations(t)
}

func TestEventController_Find_Error(t *testing.T) {
	router, repoMock, controller := setupTest(t)

	repoMock.On("Find", mock.Anything, "1").Return(nil, errors.New("error"))

	router.GET("/events/:id", controller.Find)

	request, _ := http.NewRequest("GET", "/events/1", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	repoMock.AssertExpectations(t)
}

func TestEventController_Update(t *testing.T) {
	router, repoMock, controller := setupTest(t)

	event := &entities.Event{
		ID:          "1",
		Title:       "Updated Event",
		Description: "Updated Description",
	}

	repoMock.On("Update", mock.Anything, event).Return(nil)

	router.PUT("/events/:id", controller.Update)

	body, _ := json.Marshal(event)
	request, _ := http.NewRequest("PUT", "/events/1", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	repoMock.AssertExpectations(t)
}

func TestEventController_Update_Error(t *testing.T) {
	router, repoMock, controller := setupTest(t)

	event := &entities.Event{
		ID:          "1",
		Title:       "Updated Event",
		Description: "Updated Description",
	}

	repoMock.On("Update", mock.Anything, event).Return(errors.New("error"))

	router.PUT("/events/:id", controller.Update)

	body, _ := json.Marshal(event)
	request, _ := http.NewRequest("PUT", "/events/1", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	repoMock.AssertExpectations(t)
}

func TestEventController_Delete(t *testing.T) {
	router, repoMock, controller := setupTest(t)

	repoMock.On("Delete", mock.Anything, "1").Return(nil)

	router.DELETE("/events/:id", controller.Delete)

	request, _ := http.NewRequest("DELETE", "/events/1", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	repoMock.AssertExpectations(t)
}

func TestEventController_Delete_Error(t *testing.T) {
	router, repoMock, controller := setupTest(t)

	repoMock.On("Delete", mock.Anything, "1").Return(errors.New("error"))

	router.DELETE("/events/:id", controller.Delete)

	request, _ := http.NewRequest("DELETE", "/events/1", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	repoMock.AssertExpectations(t)
}

func TestEventController_Create_BadRequest(t *testing.T) {
	router, _, controller := setupTest(t)
	router.POST("/events", controller.Create)

	// Kirim payload kosong (invalid)
	request, _ := http.NewRequest("POST", "/events", bytes.NewBuffer([]byte(`{}`)))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestEventController_Update_BadRequest(t *testing.T) {
	router, _, controller := setupTest(t)
	router.PUT("/events/:id", controller.Update)

	// Kirim payload kosong (invalid)
	request, _ := http.NewRequest("PUT", "/events/1", bytes.NewBuffer([]byte(`{}`)))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}
