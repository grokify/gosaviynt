package gosaviynt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	RelURLLogin = "/ECM/api/login"
	RelURLECM   = "/ECM"
	RelURLAPI   = "/api/v5"
)

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewClient returns a new `*http.Client` after retrieving a new OAuth access token from
// the login API.
func NewClient(baseURL, username, password string) (*http.Client, error) {
	if tok, err := FetchToken(baseURL, username, password); err != nil {
		return nil, err
	} else {
		cfg := oauth2.Config{}
		return cfg.Client(context.Background(), tok), nil
	}
}

// FetchToken retrieves a new token from the login API.
func FetchToken(baseURL, username, password string) (*oauth2.Token, error) {
	reqBody := LoginRequestBody{
		Username: username,
		Password: password}
	b, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(baseURL+RelURLLogin, "application/json; charset=utf-8", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("login api status code is (%d)", resp.StatusCode)
	}
	if b, err = io.ReadAll(resp.Body); err != nil {
		return nil, err
	} else {
		tok := &oauth2.Token{}
		return tok, json.Unmarshal(b, tok)
	}
}
