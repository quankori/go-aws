package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/quankori/go-aws/api/response"
	"github.com/quankori/go-aws/api/services"
)

// Router func
func RouteHost(g *echo.Group) {
	g.GET("/hostname", hostname)
}

// @Summary Get hostname data
// @Tags Hostname
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.EchoR
// @Failure 400 {object} response.EchoR
// @Router /hostname [get]
func hostname(c echo.Context) error {
	r := response.EchoResponse(c)
	data := services.Hostname()
	return r.OK(data)
}
