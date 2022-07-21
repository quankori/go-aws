package api

import (
	"github.com/labstack/echo/v4"
	"github.com/quankori/go-aws/api/controllers"
)

// Router func
func Router(e *echo.Echo) {

	apiV1 := e.Group("")
	controllers.RouteHost(apiV1)
	controllers.RouterS3(apiV1)
}
