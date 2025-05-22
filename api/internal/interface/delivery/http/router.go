package http

import (
	"goevents/internal/interface/delivery/http/v1/controllers"

	_ "goevents/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOption struct {
	EventController *controllers.EventController
}

// @title       GoEvents API
// @description GoEvents API
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
// @schemes     http
// @license.name MIT
func NewRouter(router *gin.Engine, opts RouterOption) *gin.Engine {
	{
		v1 := router.Group("/v1")

		v1.POST("/events", opts.EventController.Create)
		v1.GET("/events", opts.EventController.FindAll)
		v1.GET("/events/:id", opts.EventController.Find)
		v1.PUT("/events/:id", opts.EventController.Update)
		v1.DELETE("/events/:id", opts.EventController.Delete)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
