package controllers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/quankori/go-aws/api/response"
	"github.com/quankori/go-aws/api/services"
)

// Router func
func RouterS3(g *echo.Group) {
	g.POST("/s3", s3)
}

// @Summary Post upload file
// @Tags S3
// @Description Upload file
// @Accept  multipart/form-data
// @Produce  json
// @Param   file formData file true  "this is a test file"
// @Success 200 {object} response.EchoR
// @Failure 400 {object} response.EchoR
// @Router /s3 [post]
func s3(c echo.Context) error {
	file, _ := c.FormFile("file")
	fmt.Println(file)
	r := response.EchoResponse(c)
	data := services.LocalIP()
	return r.OK(data)
}
