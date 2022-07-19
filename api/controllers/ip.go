package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/quankori/go-aws/api/response"
	"github.com/quankori/go-aws/api/services"
)

// Router func
func RouterIP(g *echo.Group) {
	g.GET("/ip", ip)
}

// @Summary Get ip data
// @Tags IP
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.EchoR
// @Failure 400 {object} response.EchoR
// @Router /ip [get]
func ip(c echo.Context) error {
	r := response.EchoResponse(c)
	data := services.LocalIP()
	return r.OK(data)
}
