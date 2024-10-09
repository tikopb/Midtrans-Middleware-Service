package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/tikopb/Midtrans-Middleware-Service/internal/model"
	"io"
	"net/http"

	master "github.com/tikopb/Midtrans-Middleware-Service/internal/main-module"
)

type Repository interface {
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
	midtransBaseUrl := m.master.GetEnvVariabel("midtrans_base_url") + "v1/payment-links"
	serverKey := m.master.GetEnvVariabel("serverKey")
	method := "POST"

	// Convert body to JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		return model.MidtransPaymentLinkRespont{}, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, midtransBaseUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return model.MidtransPaymentLinkRespont{}, err
	}

	// Encode the serverKey for Basic Auth
	encodedServerKey := base64.StdEncoding.EncodeToString([]byte(serverKey + ":"))

	// Set headers
	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Basic "+encodedServerKey)
	req.Header.Add("content-type", "application/json")

	// Send request
	res, err := client.Do(req)
	if err != nil {
		return model.MidtransPaymentLinkRespont{}, err
	}
	defer res.Body.Close()

	// Check for non-2xx status codes
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		// Read response body to see the error message
		bodyBytes, _ := io.ReadAll(res.Body)
		return model.MidtransPaymentLinkRespont{}, fmt.Errorf("received non-2xx response code: %d, body: %s", res.StatusCode, string(bodyBytes))
	}

	// Read response
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return model.MidtransPaymentLinkRespont{}, err
	}

	// Unmarshal the response body into your response struct
	var response model.MidtransPaymentLinkRespont
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return model.MidtransPaymentLinkRespont{}, err
	}

	return response, nil
}
