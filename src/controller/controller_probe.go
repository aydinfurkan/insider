package controller

import (
	"insider/src/infra/api"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProbeController struct {
}

func NewProbeController() *ProbeController {
	return &ProbeController{}
}

// HealthCheck performs a health check
// @Summary Health check endpoint
// @Description Check if the server is running
// @Tags health
// @Produce json
// @Success 200 {object} api.ApiResponse{data=string} "Server is running"
// @Router /healthcheck [get]
func (p *ProbeController) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, api.NewSuccessResponse("Server is running"))
}

// PingPong responds with pong
// @Summary Ping endpoint
// @Description Simple ping endpoint that responds with pong
// @Tags health
// @Produce json
// @Success 200 {object} api.ApiResponse{data=string} "pong"
// @Router /ping [get]
func (p *ProbeController) PingPong(c echo.Context) error {
	return c.JSON(http.StatusOK, api.NewSuccessResponse("pong"))
}
