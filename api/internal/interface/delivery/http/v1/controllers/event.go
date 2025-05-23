package controllers

import (
	"goevents/internal/domain/entities"
	"goevents/internal/domain/usecases"
	"goevents/internal/interface/delivery/http/v1/requests"
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
// @Accept json
// @Produce json
// @Param event body requests.CreateEventRequest true "Event"
// @Router /events [post]
func (c *EventController) Create(ctx *gin.Context) {
	var req requests.CreateEventRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create event",
			"success": false,
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	event := &entities.Event{
		Title:       req.Title,
		Description: req.Description,
		Date:        req.Date,
	}

	err := c.uc.Create(ctx, event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create event",
			"success": false,
			"code":    http.StatusInternalServerError,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event created",
		"success": true,
		"code":    http.StatusOK,
		"data":    event,
	})
}

// @Summary Find all events
// @Description Find all events
// @Tags event
// @Router /events [get]
func (c *EventController) FindAll(ctx *gin.Context) {
	events, err := c.uc.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"success": false,
			"code":    http.StatusInternalServerError,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Fetch events success",
		"success": true,
		"code":    http.StatusOK,
		"data":    events,
	})
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
