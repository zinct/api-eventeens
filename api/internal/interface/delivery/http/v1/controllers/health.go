package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	startTime time.Time
	db        func() error
}

type HealthResponse struct {
	Status    string `json:"status"`
	Uptime    string `json:"uptime"`
	Timestamp string `json:"timestamp"`
}

func NewHealthController(db func() error) *HealthController {
	return &HealthController{
		startTime: time.Now(),
		db:        db,
	}
}

// Liveness godoc
// @Summary Check if the service is alive
// @Description Returns a 200 OK response if the service is running
// @Tags health
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health/liveness [get]
func (c *HealthController) Liveness(ctx *gin.Context) {
	response := HealthResponse{
		Status:    "UP",
		Uptime:    time.Since(c.startTime).String(),
		Timestamp: time.Now().Format(time.RFC3339),
	}

	ctx.JSON(http.StatusOK, response)
}

// Readiness godoc
// @Summary Check if the service is ready to receive traffic
// @Description Returns a 200 OK response if the service is ready to serve requests, including database connectivity
// @Tags health
// @Produce json
// @Success 200 {object} HealthResponse
// @Failure 503 {object} HealthResponse
// @Router /health/readiness [get]
func (c *HealthController) Readiness(ctx *gin.Context) {
	response := HealthResponse{
		Status:    "UP",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	// Check database connectivity
	if err := c.db(); err != nil {
		response.Status = "DOWN"
		ctx.JSON(http.StatusServiceUnavailable, response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
