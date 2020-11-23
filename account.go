package novadax

// SubAccount stands for NovaDAX's sub account resource
type SubAccount struct {
	SubID       string `json:"subId"`
	State       string `json:"state"`
	SubAccount  string `json:"subAccount"`
	SubIdentify string `json:"subIdentify"`
}

// ListAccountSubs stands for the response structure for sub accounts listing API endpoint
type ListAccountSubs struct {
	Code        string        `json:"code"`
	SubAccounts []*SubAccount `json:"data"`
	Message     string        `json:"message"`
}

// AccountSubs returns current available sub accounts for NovaDAX
func (client *Client) AccountSubs() ([]*SubAccount, error) {
	req, err := client.buildRequest("GET", "/v1/account/subs", nil, true)
	if err != nil {
		return nil, err
	}
	var response ListAccountSubs
	_, err = client.do(req, &response)
	return response.SubAccounts, err
}

// Balance represents a coin balance
type Balance struct {
	Currency  string `json:"currency"`
	Balance   string `json:"balance"`
	Hold      string `json:"hold"`
	Available string `json:"available"`
	AccountID string `json:"accountId"`
}

// GetBalance stands for the response structure for account get balance API endpoint
type GetBalance struct {
	Code     string     `json:"code"`
	Balances []*Balance `json:"data"`
	Message  string     `json:"message"`
}

// AccountGetBalance returns current account balances
func (client *Client) AccountGetBalance() ([]*Balance, error) {
	req, err := client.buildRequest("GET", "/v1/account/getBalance", nil, true)
	if err != nil {
		return nil, err
	}
	var response GetBalance
	_, err = client.do(req, &response)
	return response.Balances, err
}
