package main

import (
	"github.com/labstack/echo/v4"
	"github.com/tikopb/Midtrans-Middleware-Service/internal/delivery/logger"
	"github.com/tikopb/Midtrans-Middleware-Service/internal/delivery/rest"
	"github.com/tikopb/Midtrans-Middleware-Service/internal/delivery/service"
	"net/http"
)

func main() {
	logger.Init()

	e := echo.New()
	service := service.GetRepository()
	handler := rest.NewHandler(service)
	rest.LoadRoute(e, handler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
