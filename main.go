package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/quankori/go-aws/internals/ip"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		ip := string(ip.LocalIP())
		msg := fmt.Sprint("Server is running IP: ", ip)
		return c.JSON(http.StatusOK, msg)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
