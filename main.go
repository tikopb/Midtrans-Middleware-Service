package main

import (
	"github.com/labstack/echo/v4"
	"github.com/tikopb/Midtrans-Middleware-Service/internal/delivery/logger"
	"github.com/tikopb/Midtrans-Middleware-Service/internal/delivery/rest"
	"github.com/tikopb/Midtrans-Middleware-Service/internal/delivery/service"
	master "github.com/tikopb/Midtrans-Middleware-Service/internal/main-module"
	"net/http"
)

func main() {
	logger.Init()
	e := echo.New()

	//load main service
	repository := master.GetRepository()
	midtransService := service.GetRepository(&repository)
	handler := rest.NewHandler(midtransService)
	rest.LoadRoute(e, handler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
