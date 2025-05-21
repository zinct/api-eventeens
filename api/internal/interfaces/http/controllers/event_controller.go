package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewEventController() *EventController {
	return &EventController{}
}

type EventController struct {
}

func (c *EventController) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Event created"})
}

func (c *EventController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Event found"})
}

func (c *EventController) Find(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Event found"})
}

func (c *EventController) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func (c *EventController) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
