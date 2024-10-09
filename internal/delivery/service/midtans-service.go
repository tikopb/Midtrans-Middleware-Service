package service

import (
	"bytes"
	"encoding/json"
	"github.com/tikopb/Midtrans-Middleware-Service/internal/model"
	"io/ioutil"
	"net/http"

	master "github.com/tikopb/Midtrans-Middleware-Service/internal/main-module"
)

type Repository interface {
	//CreatePayment
	CreatePaymentLink(request model.MidtransPaymentLinkRequest) (model.MidtransPaymentLinkRespont, error)
}

type midtrans struct {
	master master.Repository
}

func GetRepository(master master.Repository) Repository {
	return &midtrans{
		master: master,
	}
}

func (m *midtrans) CreatePaymentLink(request model.MidtransPaymentLinkRequest) (model.MidtransPaymentLinkRespont, error) {
	midtrans_base_url := m.master.GetEnvVariabel("midtrans_base_url") + "v1/payment-links'"
	serverKey := m.master.GetEnvVariabel("serverKey")
	method := "POST"

	// Convert body to JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		return model.MidtransPaymentLinkRespont{}, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, midtrans_base_url, bytes.NewBuffer(jsonData))
	if err != nil {
		return model.MidtransPaymentLinkRespont{}, err
	}

	// Set headers
	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Basic "+serverKey)
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
	return response, nil
}
