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

var (
	cliendID     = ""
	clientSecret = ""
)

// CreateAuthorizeURL : 認証用URLを作成
func CreateAuthorizeURL() (string, error) {
	url, err := url.Parse(authorizeURL)
	if err != nil {
		return "", err
	}

	q := url.Query()
	q.Add("client_id", cliendID)
	q.Add("response_type", "code")
	q.Add("redirect_uri", redirectURL)
	q.Add("scope", "read write")
	url.RawQuery = q.Encode()

	return url.String(), nil
}

// FetchToken : トークンを取得
func FetchToken(code string) (*Credencial, error) {
	req, err := http.NewRequest(http.MethodPost, tokenURL, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("client_id", cliendID)
	q.Add("client_secret", clientSecret)
	q.Add("grant_type", "authorization_code")
	q.Add("redirect_uri", redirectURL)
	q.Add("code", strings.TrimSpace(code))
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	decorder := json.NewDecoder(res.Body)

	cred := &Credencial{}
	if err := decorder.Decode(cred); err != nil {
		return nil, err
	}

	return cred, nil
}
