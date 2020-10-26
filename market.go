package novadax

// MarketTicker stands for the NovaDAX API market ticker resource
type MarketTicker struct {
	Ask            string `json:"ask"`
	BaseVolume24h  string `json:"baseVolume24h"`
	Bid            string `json:"bid"`
	High24h        string `json:"high24h"`
	LastPrice      string `json:"lastPrice"`
	Low24h         string `json:"low24h"`
	Open24h        string `json:"open24h"`
	QuoteVolume24h string `json:"quoteVolume24h"`
	Symbol         string `json:"symbol"`
	Timestamp      int64  `json:"timestamp"`
}

// GetLatestMarketTickersResponse stands for the response structure for latest market tickers API endpoint
type GetLatestMarketTickersResponse struct {
	Code          string          `json:"code"`
	MarketTickers []*MarketTicker `json:"data"`
	Message       string          `json:"message"`
}

// GetMarketTickerResponse stands for the response structure for market ticker API endpoint
type GetMarketTickerResponse struct {
	Code         string        `json:"code"`
	MarketTicker *MarketTicker `json:"data"`
	Message      string        `json:"message"`
}

// GetMarketTickersFilters stands for the GetLatestTickers possible and required filters
type GetMarketTickersFilters struct {
	Symbol string `json:"symbol"`
}

// GetLatestTickers returns latest market tickers for all key pairs in NovaDAX
func (client *Client) GetLatestTickers(filters *GetMarketTickersFilters) ([]*MarketTicker, error) {
	params := structToURLValues(filters)

	path := "/v1/market/tickers"
	if params.Encode() != "" {
		path += "?" + params.Encode()
	}

	req, err := client.buildRequest("GET", path, nil, false)
	if err != nil {
		return nil, err
	}

	var response GetLatestMarketTickersResponse
	_, err = client.do(req, &response)
	return response.MarketTickers, err
}

// GetMarketTicker returns latest market tickers for a single key pairs in NovaDAX
func (client *Client) GetMarketTicker(filters *GetMarketTickersFilters) (*MarketTicker, error) {
	params := structToURLValues(filters)

	path := "/v1/market/ticker"
	if params.Encode() != "" {
		path += "?" + params.Encode()
	}

	req, err := client.buildRequest("GET", path, nil, false)
	if err != nil {
		return nil, err
	}

	var response GetMarketTickerResponse
	_, err = client.do(req, &response)
	return response.MarketTicker, err
}

// MarketDepth stands for the NovaDAX API market depth resource
type MarketDepth struct {
	Asks      [][]string `json:"asks"` // 0 for price and 1 for amount
	Bids      [][]string `json:"bids"` // 0 for price and 1 for amount
	Timestamp int64      `json:"timestamp"`
}

// GetMarketDepthResponse stands for the response structure for market depth API endpoint
type GetMarketDepthResponse struct {
	Code        string       `json:"code"`
	MarketDepth *MarketDepth `json:"data"`
	Message     string       `json:"message"`
}

// GetMarketDepthFilters stands for the GetLatestTickers possible and required filters
type GetMarketDepthFilters struct {
	Symbol string `json:"symbol"`
	Limit  int    `json:"limit"`
}

// GetMarketDepth returns limit orders for a key pair in NovaDAX
func (client *Client) GetMarketDepth(filters *GetMarketDepthFilters) (*MarketDepth, error) {
	params := structToURLValues(filters)

	path := "/v1/market/depth"
	if params.Encode() != "" {
		path += "?" + params.Encode()
	}

	req, err := client.buildRequest("GET", path, nil, false)
	if err != nil {
		return nil, err
	}

	var response GetMarketDepthResponse
	_, err = client.do(req, &response)
	return response.MarketDepth, err
}
