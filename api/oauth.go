package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURL      = "https://api.annict.com"
	authorizeURL = "https://api.annict.com/oauth/authorize"
	tokenURL     = "https://api.annict.com/oauth/token"
	revokeURL    = "https://api.annict.com/oauth/revoke"
	redirectURL  = "urn:ietf:wg:oauth:2.0:oob"
)

// CreateAuthorizeURL : 認証用URLを作成
func (c *Client) CreateAuthorizeURL() (string, error) {
	url, err := url.Parse(authorizeURL)
	if err != nil {
		return "", err
	}

	q := url.Query()
	q.Add("client_id", c.Token.Client.ID)
	q.Add("response_type", "code")
	q.Add("redirect_uri", redirectURL)
	q.Add("scope", "read write")
	url.RawQuery = q.Encode()

	return url.String(), nil
}

// UpdateUserToken : ユーザートークンを更新 (再取得)
func (c *Client) UpdateUserToken(code string) error {
	req, err := http.NewRequest(http.MethodPost, tokenURL, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("client_id", c.Token.Client.ID)
	q.Add("client_secret", c.Token.Client.Secret)
	q.Add("grant_type", "authorization_code")
	q.Add("redirect_uri", redirectURL)
	q.Add("code", strings.TrimSpace(code))
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	decorder := json.NewDecoder(res.Body)

	userToken := &UserToken{}
	if err := decorder.Decode(userToken); err != nil {
		return err
	}

	c.Token.User = userToken
	return nil
}
