package coinbase

import (
	"github.com/BillotP/coinbase/lib/auth"
	"github.com/BillotP/coinbase/lib/models"
	"github.com/BillotP/coinbase/lib/rpc"
)

// Client is the struct from which all API requests are made
type Client struct {
	rpc rpc.RPC
}

// CoinbaseClient is the global blablabla
var CoinbaseClient *Client

// New return an authenticated client
func New(pubkey *string, privkey *string) *Client {
	auth.CBAccount = auth.NewClient(pubkey, privkey)
	CoinbaseClient = &Client{
		rpc: rpc.RPC{
			Auth: auth.CBAccount,
		},
	}
	return CoinbaseClient
}

// Get sends a GET request and marshals response data into holder
func (c Client) Get(path string, params interface{}, holder interface{}) error {
	return c.rpc.Request("GET", path, params, &holder)
}

// Post sends a POST request and marshals response data into holder
func (c Client) Post(path string, params interface{}, holder interface{}) error {
	return c.rpc.Request("POST", path, params, &holder)
}

// Delete sends a DELETE request and marshals response data into holder
func (c Client) Delete(path string, params interface{}, holder interface{}) error {
	return c.rpc.Request("DELETE", path, params, &holder)
}

// Put sends a PUT request and marshals response data into holder
func (c Client) Put(path string, params interface{}, holder interface{}) error {
	return c.rpc.Request("PUT", path, params, &holder)
}

// GetAccounts returns a list of all coinbase accounts
func (c Client) GetAccounts() (*models.Accounts, error) {
	var accounts models.Accounts
	if err := c.Get("v2/accounts", nil, &accounts); err != nil {
		return nil, err
	}
	return &accounts, nil
}

// GetAccountByID returns a coinbase account by its ID
func (c Client) GetAccountByID(accountID string) (*models.Accounts, error) {
	var accounts models.Accounts
	if err := c.Get("v2/accounts/"+accountID, nil, &accounts); err != nil {
		return nil, err
	}
	return &accounts, nil
}
