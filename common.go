package novadax

// Symbol stands for NOVADAX's available symbol
type Symbol struct {
	Symbol          string `json:"symbol"`
	BaseCurrency    string `json:"baseCurrency"`
	QuoteCurrency   string `json:"quoteCurrency"`
	AmountPrecision int    `json:"amountPrecision"`
	PricePrecision  int    `json:"pricePrecision"`
	ValuePrecision  int    `json:"valuePrecision"`
	MinOrderAmount  string `json:"minOrderAmount"`
	MinOrderValue   string `json:"minOrderValue"`
}

// ListSymbolsResponse stands for the response structure for symbol listing API endpoint
type ListSymbolsResponse struct {
	Code    string    `json:"code"`
	Symbols []*Symbol `json:"data"`
}

// ListSymbols returns current available symbols for NOVADAX
func (client *Client) ListSymbols() ([]*Symbol, error) {
	req, err := client.buildRequest("GET", "/v1/common/symbols", nil, false)
	if err != nil {
		return nil, err
	}
	var response ListSymbolsResponse
	_, err = client.do(req, &response)
	return response.Symbols, err
}
