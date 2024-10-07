package rest

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tikopb/Midtrans-Middleware-Service/internal/model"
)

func (h *handler) CreatePaymentLink(c echo.Context) error {
	//get query filename
	request := model.MidtransPaymentLinkRequest{}
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		WriteLogErorr("[delivery][rest][payment_handler][CreatePaymentLink] ", err)
		return err
	}

	data, err := h.service.CreatePaymentLink(request)
	if err != nil {
		return err
	}

	fmt.Println(data)
	return nil
}
