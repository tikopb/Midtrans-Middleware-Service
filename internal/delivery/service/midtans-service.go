package service

import (
	"bytes"
	"encoding/json"
	"github.com/tikopb/Midtrans-Middleware-Service/internal/model"
	"io/ioutil"
	"net/http"
)

type Repository interface {
	//CreatePayment
	CreatePaymentLink(request model.MidtransPaymentLinkRequest) (model.MidtransPaymentLinkRespont, error)
}

type midtrans struct {
}

func GetRepository() Repository {
	return &midtrans{}
}

func (m *midtrans) CreatePaymentLink(request model.MidtransPaymentLinkRequest) (model.MidtransPaymentLinkRespont, error) {
	url := "https://api.sandbox.midtrans.com/v1/payment-links"
	method := "POST"

	// Convert body to JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		return model.MidtransPaymentLinkRespont{}, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return model.MidtransPaymentLinkRespont{}, err
	}

	// Set headers
	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Basic U0ItTWlkLXNlcnZlci1VX2ZaYzhtT0FteW9zSk83cm9vNzN6Yl86")
	req.Header.Add("content-type", "application/json")

	// Send request
	res, err := client.Do(req)
	if err != nil {
		return model.MidtransPaymentLinkRespont{}, err
	}
	defer res.Body.Close()

	// Read response
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.MidtransPaymentLinkRespont{}, err
	}

	// Parse the response into the MidtransPaymentLinkRespont struct
	var response model.MidtransPaymentLinkRespont
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return model.MidtransPaymentLinkRespont{}, err
	}

	//TODO implement mef
	panic("implement me")
}
