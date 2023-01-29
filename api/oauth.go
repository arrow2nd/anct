package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

const (
	// AuthURL : 認証
	AuthURL = "https://api.annict.com/oauth/authorize"
	// TokenURL : トークン発行
	TokenURL = "https://api.annict.com/oauth/token"
	// RevokeURL : トークン破棄
	RevokeURL = "https://api.annict.com/oauth/revoke"
	// RedirectURL : リダイレクト
	RedirectURL = "urn:ietf:wg:oauth:2.0:oob"
)

var (
	// ClientID : クライアントID (ビルド時に埋め込み)
	ClientID = ""
	// ClientSecret : クライアントシークレット (ビルド時に埋め込み)
	ClientSecret = ""
)

// Credencial : 認証情報
type Credencial struct {
	// AccessToken : トークン
	AccessToken string `json:"access_token"`
}

// GetAuthorizeURL : 認証用URLを取得
func GetAuthorizeURL() (string, error) {
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

// GetToken : トークンを取得
func GetToken(code string) (*Credencial, error) {
	req, err := http.NewRequest(http.MethodPost, TokenURL, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("client_id", ClientID)
	q.Add("client_secret", ClientSecret)
	q.Add("grant_type", "authorization_code")
	q.Add("redirect_uri", RedirectURL)
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

	raw := &Credencial{}
	if err := decorder.Decode(raw); err != nil {
		return nil, err
	}

	return raw, nil
}
