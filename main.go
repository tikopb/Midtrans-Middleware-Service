package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tikopb/Midtrans-Middleware-Service/internal/delivery/logger"
	"net/http"
	"os"
)

func main() {
	logger.Init()

	home := os.Getenv("HOME_app")
	fmt.Println("", home)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
