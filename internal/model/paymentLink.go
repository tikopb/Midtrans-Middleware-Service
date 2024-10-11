package model

type MidtransPaymentLinkRespont struct {
	OrderId    string   `json:"order_id"`
	PaymentUrl string   `json:"payment_url"`
	Message    []string `json:"error_messages"`
}

// Define the MidtransPaymentLinkRequest struct that combines the others
type MidtransPaymentLinkRequest struct {
	TransactionDetails TransactionDetails `json:"transaction_details"`
	ItemDetails        []ItemDetails      `json:"item_details"`
	CustomerDetails    CustomerDetails    `json:"customer_details"`
	UsageLimit         int                `json:"usage_limit"`
}

// Define the TransactionDetails struct
type TransactionDetails struct {
	OrderId       string `json:"order_id"`
	GrossAmount   int    `json:"gross_amount"`
	PaymentLinkId string `json:"payment_link_id"`
}

// Define the ItemDetails struct
type ItemDetails struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Quantity     int    `json:"quantity"`
	Brand        string `json:"brand"`
	Category     string `json:"category"`
	MerchantName string `json:"merchant_name"`
}

// Define the CustomerDetails struct
type CustomerDetails struct {
	FirstName                     string   `json:"first_name"`
	LastName                      string   `json:"last_name"`
	Phone                         string   `json:"phone"`
	Notes                         string   `json:"notes"`
	CustomerDetailsRequiredFields []string `json:"customer_details_required_fields"`
}

//------

type MidtransPaymentStatus struct {
	OrderId   string `json:"order_id"`
	Purchases []struct {
		PaymentStatus string `json:"payment_status"`
		PaymentMethod string `json:"payment_method"`
	} `json:"purchases"`
}

type MidtransPaymentStatusResponse struct {
	OrderId       string `json:"order_id"`
	PaymentStatus string `json:"payment_status"`
	PaymentMethod string `json:"payment_method"`
}
