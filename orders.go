package novadax

import (
	"errors"
	"fmt"
)

// Order stands for the order data to be send to NovaDAX API
type Order struct {
	Symbol    string `json:"symbol"`
	Type      string `json:"type"`
	Side      string `json:"side"`
	Price     string `json:"price,omitempty"`
	Amount    string `json:"amount"`
	AccountID string `json:"accountId,omitempty"`
}

// OrderDetails stands for the order data returned by NovaDAX API
type OrderDetails struct {
	ID           string `json:"id"`
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	Side         string `json:"side"`
	Price        string `json:"price"`
	AveragePrice string `json:"averagePrice"`
	Amount       string `json:"amount"`
	FilledAmount string `json:"filledAmount"`
	Value        string `json:"value"`
	FilledValue  string `json:"filledValue"`
	FilledFee    string `json:"filledFee"`
	Status       string `json:"status"`
	Timestamp    int64  `json:"timestamp"`
}

// OrderDetailsResponse returns current orders in NovaDAX
type OrderDetailsResponse struct {
	Code         string        `json:"code"`
	OrderDetails *OrderDetails `json:"data"`
}

// OrdersDetailsResponse returns current orders in NovaDAX
type OrdersDetailsResponse struct {
	Code          string          `json:"code"`
	OrdersDetails []*OrderDetails `json:"data"`
}

// ListOrdersFilters stands for the ListOrders possible and required filters
type ListOrdersFilters struct {
	Symbol        string `json:"symbol"`
	Status        string `json:"status"`
	FromID        string `json:"fromId"`
	ToID          string `json:"toId"`
	FromTimestamp int64  `json:"fromTimestamp"`
	ToTimestamp   int64  `json:"toTimestamp"`
	Limit         int    `json:"limit"`
}

type GetOrderDetailsFilters struct {
	ID string `json:"id"`
}

func (client *Client) GetOrderDetails(ID string) (*OrderDetails, error) {
	if ID == "" {
		return nil, errors.New("ID filter is required")
	}

	params := structToURLValues(&GetOrderDetailsFilters{
		ID: ID,
	})

	path := "/v1/orders/get"
	if params.Encode() != "" {
		path += "?" + params.Encode()
	}

	req, err := client.buildRequest("GET", path, nil, true)
	if err != nil {
		return nil, err
	}

	var orderDetailsResponse OrderDetailsResponse
	_, err = client.do(req, &orderDetailsResponse)
	if err != nil {
		return nil, err
	}

	return orderDetailsResponse.OrderDetails, err
}

// ListOrders returns current market orders based on filters
func (client *Client) ListOrders(filters *ListOrdersFilters) ([]*OrderDetails, error) {
	if filters.Symbol == "" {
		return nil, errors.New("Symbol filter is required")
	}

	params := structToURLValues(filters)

	path := "/v1/orders/list"
	if params.Encode() != "" {
		path += "?" + params.Encode()
	}

	req, err := client.buildRequest("GET", path, nil, true)
	if err != nil {
		return nil, err
	}

	var ordersDetailsResponse OrdersDetailsResponse
	_, err = client.do(req, &ordersDetailsResponse)
	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err.Error())
	}

	return ordersDetailsResponse.OrdersDetails, err
}

// CreateOrder creates a new order and return it's details
func (client *Client) CreateOrder(order *Order) (*OrderDetails, error) {
	if order.Symbol == "" || order.Type == "" || order.Amount == "" || order.Side == "" {
		return nil, errors.New("Missing required fields")
	}

	if order.Type == "LIMIT" && order.Price == "" {
		return nil, errors.New("Price is required for limit orders")
	}

	req, err := client.buildRequest("POST", "/v1/orders/create", order, true)
	if err != nil {
		return nil, err
	}

	var orderDetailsResponse OrderDetailsResponse
	_, err = client.do(req, &orderDetailsResponse)
	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err.Error())
	}

	return orderDetailsResponse.OrderDetails, err
}

type CancelResult struct {
	Success bool `json:"result"`
}

type CancelResponse struct {
	Code    string       `json:"code"`
	Result  CancelResult `json:"data"`
	Message string       `json:"message"`
}

func (client *Client) CancelOrder(ID string) (bool, error) {
	if ID == "" {
		return false, errors.New("ID filter is required")
	}

	req, err := client.buildRequest("POST", "/v1/orders/cancel", &GetOrderDetailsFilters{
		ID: ID,
	}, true)
	if err != nil {
		return false, err
	}

	var cancelResponse CancelResponse
	_, err = client.do(req, &cancelResponse)
	if err != nil {
		return false, err
	}

	return cancelResponse.Result.Success, err
}
