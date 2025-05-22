package controllers

import (
	"goevents/internal/domain/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	uc usecases.EventUsecase
}

func NewEventController(usecase usecases.EventUsecase) *EventController {
	return &EventController{uc: usecase}
}

// @Summary Create a new event
// @Description Create a new event
// @Tags event
// @Router /events [post]
func (c *EventController) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Event created"})
}

// @Summary Find all events
// @Description Find all events
// @Tags event
// @Router /events [get]
func (c *EventController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Event found"})
}

// @Summary Find an event
// @Description Find an event
// @Tags event
// @Router /events/{id} [get]
func (c *EventController) Find(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Event found"})
}

// @Summary Update an event
// @Description Update an event
// @Tags event
// @Router /events/{id} [put]
func (c *EventController) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

// @Summary Delete an event
// @Description Delete an event
// @Tags event
// @Router /events/{id} [delete]
func (c *EventController) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
