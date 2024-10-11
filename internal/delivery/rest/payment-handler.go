package rest

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/tikopb/Midtrans-Middleware-Service/internal/model"
	"net/http"
)

func (h *handler) CreatePaymentLink(c echo.Context) error {
	//get query filename
	request := model.MidtransPaymentLinkRequest{}
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		WriteLogErorr("[delivery][rest][payment_handler][CreatePaymentLink] ", err)
		return handleError(c, http.StatusInternalServerError, err, model.MidtransPaymentLinkRespont{})
	}

	data, err := h.service.CreatePaymentLink(request)
	if err != nil {
		return handleError(c, http.StatusInternalServerError, err, model.MidtransPaymentLinkRespont{})
	}

	return handleError(c, http.StatusInternalServerError, errors.New("success created= "+data.OrderId), data)
}

func (h *handler) GetPaymentLink(c echo.Context) error {
	//get param
	ID := c.QueryParam("id")

	//run function
	data, err := h.service.CheckStatusPaymentLink(ID)
	if err != nil {
		WriteLogErorr("[delivery][rest][midtrans-service][GetPaymentLink] ", err)
		return handleError(c, http.StatusInternalServerError, err, model.MidtransPaymentLinkRespont{})
	}

	return handleError(c, http.StatusOK, errors.New(""), data)
}
