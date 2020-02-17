package novadax

// MarketTicker stands for the NovaDAX API market ticker resource
type MarketTicker struct {
	Ask string `json:"ask"`
	BaseVolume24h string `json:"baseVolume24h"`
	Bid	string `json:"bid"`
	High24h string `json:"high24h"`
	LastPrice string `json:"lastPrice"`
	Low24h string `json:"low24h"`
	Open24h string `json:"open24h"`
	QuoteVolume24h string `json:"quoteVolume24h"`
	Symbol string `json:"symbol"`
	Timestamp int64 `json:"timestamp"`
}

// GetLatestMarketTickersResponse stands for the response structure for latest market tickers API endpoint
type GetLatestMarketTickersResponse struct {
	Code string `json:"code"`
	MarketTickers []*MarketTicker `json:"data"`
	Message string `json:"message"`
}

// GetLatestTickers returns latest market tickers for all key pairs in NovaDAX
func (client *Client) GetLatestTickers() ([]*MarketTicker, error) {
	req, err := client.buildRequest("GET", "/v1/market/tickers", nil, false)
	if err != nil {
		return nil, err
	}
	var response GetLatestMarketTickersResponse
	_, err = client.do(req, &response)
	return response.MarketTickers, err
}
