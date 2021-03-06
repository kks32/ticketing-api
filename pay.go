package tickets

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

// Type Payment datastructure for a credit card
type Payment struct {
	// Type (card)
	Type string `json:"type, omitempty"`
	// Name of the card holder
	Name string `json:"name, omitempty"`
	// Month
	Month string `json:"expiryMonth, omitempty"`
	// Year
	Year string `json:"expiryYonth, omitempty"`
	// Card Number
	CardNumber string `json:"cardNumber, omitempty"`
	// CVV
	Cvc string `json:"cvc, omitempty"`
	// Issue Number
	IssueNumber string `json:"issueNumber, omitempty"`
	// OrderType
	OrderType string `json:"orderType, omitempty"`
	// Order description
	OrderDescription string `json:"orderDescription, omitempty"`
	// Amount
	Amount float64 `json:"amount, omitempty"`
	// CurrencyCode description
	CurrencyCode string `json:"currencyCode, omitempty"`
}

func makePayment(payment *Payment) (response *http.Response) {
	// WorldPay URL
	url := "https://api.worldpay.com/v1/orders"
	// Payment string
	var paymentString = `{"paymentMethod":{ "type":` + `"Card", "name":"` + payment.Name + `", "expiryMonth":"` + payment.Month + `", "expiryYear":"` + payment.Year + `", "cardNumber":"` + payment.CardNumber + `", "cvc":"` + payment.Cvc + `", "issueNumber":"" }, "orderType": "ECOM", "orderDescription": "` + payment.OrderDescription + `", "amount": "`  + fmt.Sprintf("%.0f",payment.Amount * 100.) + `", "currencyCode": "GBP"}`

	var jsonStr = []byte(paymentString)

	// Post request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// Authorization code
	req.Header.Set("Authorization", os.Getenv("WorldPay"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	response = resp
	return response
}

