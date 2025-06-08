package integrationtest_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goevents/internal/domain/entities"
	"goevents/internal/interface/delivery/http/v1/controllers"
	"goevents/internal/interface/delivery/http/v1/requests"
	"goevents/internal/wire"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupIntegrationTest(t *testing.T) (*gin.Engine, *controllers.EventController, sqlmock.Sqlmock) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}

	controller := wire.InitializeEventController(db)
	return router, controller, mock
}

func TestEventController_Integration_Create(t *testing.T) {
	router, controller, mock := setupIntegrationTest(t)

	req := requests.CreateEventRequest{
		Title:       "Integration Test Event",
		Description: "Integration Test Description",
		Date:        time.Now(),
	}

	mock.ExpectExec("INSERT INTO events").
		WithArgs(req.Title, req.Description, req.Date).
		WillReturnResult(sqlmock.NewResult(1, 1))

	router.POST("/events", controller.Create)

	body, _ := json.Marshal(req)
	request, _ := http.NewRequest("POST", "/events", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	fmt.Println("response", response)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestEventController_Integration_FindAll(t *testing.T) {
	router, controller, _ := setupIntegrationTest(t)

	router.GET("/events", controller.FindAll)

	request, _ := http.NewRequest("GET", "/events", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	fmt.Println("response", response)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestEventController_Integration_Find(t *testing.T) {
	router, controller, _ := setupIntegrationTest(t)

	router.GET("/events/:id", controller.Find)

	request, _ := http.NewRequest("GET", "/events/1", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestEventController_Integration_Update(t *testing.T) {
	router, controller, _ := setupIntegrationTest(t)

	event := &entities.Event{
		ID:          "1",
		Title:       "Updated Integration Event",
		Description: "Updated Integration Description",
	}

	router.PUT("/events/:id", controller.Update)

	body, _ := json.Marshal(event)
	request, _ := http.NewRequest("PUT", "/events/1", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestEventController_Integration_Delete(t *testing.T) {
	router, controller, _ := setupIntegrationTest(t)

	router.DELETE("/events/:id", controller.Delete)

	request, _ := http.NewRequest("DELETE", "/events/1", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}
