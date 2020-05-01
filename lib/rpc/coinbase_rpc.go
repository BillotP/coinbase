package rpc

import (
	"bytes"
	"coinbase/lib/auth"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// RPC handles the remote procedure call requests
type RPC struct {
	Auth *auth.APIKeyAuthentication
}

// Request sends a request with params marshaled into a JSON payload in the body
// The response value is marshaled from JSON into the specified holder struct
func (r RPC) Request(method string, endpoint string, params interface{}, holder interface{}) error {
	var err error
	var data []byte
	var req *http.Request
	var jsonParams []byte

	if params != nil {
		if jsonParams, err = json.Marshal(params); err != nil {
			return fmt.Errorf("Error(RPC.Request1): %s", err.Error())
		}
	}
	if req, err = r.CreateRequest(method, endpoint, jsonParams); err != nil {
		return fmt.Errorf("Error(RPC.Request2): %s", err.Error())
	}
	if data, err = r.ExecuteRequest(req); err != nil {
		return fmt.Errorf("Error(RPC.Request3): %s", err.Error())
	}
	if err := json.Unmarshal(data, &holder); err != nil {
		return fmt.Errorf("Error(RPC.Request4): %s", err.Error())
	}

	return nil
}

// CreateRequest formats a request with all the necessary headers
func (r RPC) CreateRequest(method string, endpoint string, params []byte) (*http.Request, error) {
	var err error
	var req *http.Request
	reqURL := r.Auth.GetBaseURL() + endpoint

	if req, err = http.NewRequest(method, reqURL, bytes.NewBuffer(params)); err != nil {
		return nil, err
	}
	// Authenticate the request
	r.Auth.Authenticate(req, endpoint, params)
	req.Header.Set("User-Agent", "CoinbaseGo/v2")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// ExecuteRequest takes a prepared http.Request and returns the body of the response
// If the response is not of HTTP Code 200, an error is returned
func (r RPC) ExecuteRequest(req *http.Request) ([]byte, error) {
	var err error
	var resp *http.Response
	if resp, err = r.Auth.GetClient().Do(req); err != nil {
		fmt.Printf("Error(RPC.ExecuteRequest1): %s\n", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bytes := buf.Bytes()
	if resp.StatusCode != 200 {
		if len(bytes) == 0 { // Log response body for debugging purposes
			log.Printf("Response body was empty")
		} else {
			log.Printf("Response body:\n\t%s\n", bytes)
			log.Printf("Headers : ")
			if reqHeadersBytes, err := json.Marshal(req.Header); err != nil {
				log.Println("Could not Marshal Req Headers")
			} else {
				fmt.Printf("%s\n", string(reqHeadersBytes))
			}
		}
		return nil, fmt.Errorf("%s %s failed. Response code was %s", req.Method, req.URL, resp.Status)
	}
	return bytes, nil
}
