package http

import (
	"goevents/internal/interfaces/http/controllers"

	"github.com/gin-gonic/gin"
)

type RouterOption struct {
	EventController *controllers.EventController
}

func NewRouter(router *gin.Engine, opts RouterOption) *gin.Engine {
	// Event
	router.POST("/events", opts.EventController.Create)
	router.GET("/events", opts.EventController.FindAll)
	router.GET("/events/:id", opts.EventController.Find)
	router.PUT("/events/:id", opts.EventController.Update)
	router.DELETE("/events/:id", opts.EventController.Delete)

	return router
}
