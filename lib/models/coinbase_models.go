package models

import "strconv"

// Balance is a coinbase balance model
type Balance struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// Currency is the coinbase model a currency
type Currency struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Colors       string `json:"color"`
	SortIndex    int64  `json:"sort_index"`
	Exponent     int64  `json:"exponent"`
	Type         string `json:"type"`
	AddressRegex string `json:"address_regex"`
	AssetID      string `json:"asset_id"`
	Slug         string `json:"slug"`
}

// Account is a coinbase account model
type Account struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Primary          bool   `json:"primary"`
	Type             string `json:"type"`
	Currency         `json:"currency"`
	Balance          `json:"balance"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	Resource         string  `json:"resource"`
	ResourcePath     string  `json:"resource_path"`
	AllowDeposit     bool    `json:"allow_deposit"`
	AllowWithdrawals bool    `json:"allow_withdrawals"`
	NativeBalance    Balance `json:"native_balance"`
}

// Pagination is the pagination struct for a Coinbase API response
type Pagination struct {
	EndingBefore         string `json:"ending_before"`
	StartingAfter        string `json:"starting_after"`
	PreviousEndingBefore string `json:"previous_ending_before"`
	NextStartingAfter    string `json:"next_starting_after"`
	Limit                int64  `json:"limit"`
	Order                string `json:"order"`
	PreviousURI          string `json:"previous_uri"`
	NextURI              string `json:"next_uri"`
}

// Response is the global response object for sucessfull Coinbase api call
type Response struct {
	Pagination `json:"pagination"`
}

// Accounts is the response model for GET `accounts` data
type Accounts struct {
	Response
	Datas []Account `json:"data"`
}

// FilterEmpty remove the empty balances from Accounts.Datas object
func (a *Accounts) FilterEmpty() {
	var fltr []Account
	empty := float64(0)
	for i := range a.Datas {
		v, _ := strconv.ParseFloat(a.Datas[i].Amount, 64)
		if v > empty {
			fltr = append(fltr, a.Datas[i])
		}
	}
}
