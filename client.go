package gosaviynt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

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
	if b, err := json.Marshal(reqBody); err != nil {
		return nil, err
	} else if resp, err := http.Post(
		strings.TrimRight(strings.TrimSpace(baseURL), "/")+RelURLLogin,
		"application/json; charset=utf-8", bytes.NewBuffer(b)); err != nil {
		return nil, err
	} else if resp.StatusCode >= 300 {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("login api status code is (%d) with body (%s)", resp.StatusCode, string(b))
	} else if b, err = io.ReadAll(resp.Body); err != nil {
		return nil, err
	} else {
		tok := &oauth2.Token{}
		return tok, json.Unmarshal(b, tok)
	}
}

func Pointer[E any](e E) *E { return &e }
