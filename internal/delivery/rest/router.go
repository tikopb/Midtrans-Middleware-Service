package rest

import "github.com/labstack/echo/v4"

func LoadRoute(e *echo.Echo, handler *handler) {
	//paymentLink
	paymentLinkGroup := e.Group("/Payment-link")
	paymentLinkGroup.POST("/", handler.CreatePaymentLink)
}
