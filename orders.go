package novadax

import (
	"errors"
	"fmt"
	"strconv"
)

// Order stands for the order data returned by NOVADAX API
type Order struct {
	ID           string  `json:"id"`
	Symbol       string  `json:"symbol"`
	Type         string  `json:"type"`
	Side         string  `json:"side"`
	Price        string  `json:"price"`
	AveragePrice string  `json:"averagePrice"`
	Amount       float64 `json:"amount"`
	FilledAmount string  `json:"filledAmount"`
	Value        string  `json:"value"`
	FilledValue  string  `json:"filledValue"`
	FilledFee    string  `json:"filledFee"`
	Status       string  `json:"status"`
	Timestamp    int64   `json:"timestamp"`
}

// ListOrdersResponse returns current orders in NOVADAX
type ListOrdersResponse struct {
	Code   string   `json:"code"`
	Orders []*Order `json:"data"`
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

// ListOrders returns current market orders based on filters
func (client *Client) ListOrders(filters *ListOrdersFilters) ([]*Order, error) {
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

	var listOrdersResponse ListOrdersResponse
	resp, err := client.do(req, &listOrdersResponse)
	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err.Error())
	}

	// test for response status code
	status, err := strconv.Atoi(resp.Status)
	if err != nil || status < 200 || status > 299 {
		return nil, fmt.Errorf("request failed with status %s", resp.Status)
	}

	return listOrdersResponse.Orders, err
}
