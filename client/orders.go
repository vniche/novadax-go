package client

import (
	"time"
)

// Order stands for the order data returned by NOVADAX API
type Order struct {
	ID           string    `json:"id"`
	Symbol       string    `json:"symbol"`
	Type         string    `json:"type"`
	Side         string    `json:"side"`
	Price        string    `json:"price"`
	AveragePrice float64   `json:"averagePrice"`
	Amount       float64   `json:"amount"`
	FilledAmount int       `json:"filledAmount"`
	Value        float64   `json:"value`
	FilledValue  float64   `json:"filledValue"`
	FilledFee    float64   `json:"filledFee"`
	Status       string    `json:"status"`
	Timestamp    time.Time `json:"timestamp"`
}

// ListOrdersResponse returns current orders in NOVADAX
type ListOrdersResponse struct {
	Code   string   `json:"code"`
	Orders []*Order `json:"data"`
}

// ListOrders returns current market orders based on filters
func (client *Client) ListOrders() ([]*Order, error) {
	req, err := client.signRequest("GET", "/v1/orders/list", nil, true)
	if err != nil {
		return nil, err
	}
	var response ListOrdersResponse
	_, err = client.do(req, &response)
	return response.Orders, err
}
