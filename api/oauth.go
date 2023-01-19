package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/manifoldco/promptui"
)

const (
	AuthURL     = "https://api.annict.com/oauth/authorize"
	TokenURL    = "https://api.annict.com/oauth/token"
	RevokeURL   = "https://api.annict.com/oauth/revoke"
	RedirectURL = "urn:ietf:wg:oauth:2.0:oob"
)

var (
	ClientID     = ""
	ClientSecret = ""
)

func Authorize() error {
	url, err := createAuthorizeURL()
	if err != nil {
		return err
	}

	fmt.Printf("Auth URL: %s\n", url)

	code, err := inputAuthCode()
	if err != nil {
		return err
	}

	result, err := getToken(strings.TrimSpace(code))
	if err != nil {
		return err
	}

	log.Printf("Token: %s\n", result.AccessToken)
	log.Printf("Type: %s\n", result.TokenType)

	return nil
}

func createAuthorizeURL() (string, error) {
	url, err := url.Parse(AuthURL)
	if err != nil {
		return "", err
	}

	q := url.Query()
	q.Add("client_id", ClientID)
	q.Add("response_type", "code")
	q.Add("redirect_uri", RedirectURL)
	q.Add("scope", "read write")
	url.RawQuery = q.Encode()

	return url.String(), nil
}

func inputAuthCode() (string, error) {
	prompt := promptui.Prompt{
		Label: "Code",
		Validate: func(s string) error {
			if s == "" {
				return errors.New("please enter a code")
			}
			return nil
		},
	}

	return prompt.Run()
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func getToken(code string) (*TokenResponse, error) {
	req, err := http.NewRequest(http.MethodPost, TokenURL, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("client_id", ClientID)
	q.Add("client_secret", ClientSecret)
	q.Add("grant_type", "authorization_code")
	q.Add("redirect_uri", RedirectURL)
	q.Add("code", code)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	decorder := json.NewDecoder(res.Body)

	raw := &TokenResponse{}
	if err := decorder.Decode(raw); err != nil {
		return nil, err
	}

	return raw, nil
}
