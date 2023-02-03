package api

import (
	"fmt"
	"net/http"
)

var (
	clientID     = ""
	clientSecret = ""
)

// ClientToken : クライアントトークン
type ClientToken struct {
	ID     string
	Secret string
}

// UserToken : ユーザートークン
type UserToken struct {
	Bearer string `json:"access_token"`
}

// Token : トークン
type Token struct {
	Client *ClientToken
	User   *UserToken
}

// Revoke : ユーザートークンを失効させる
func (t *Token) Revoke() error {
	req, err := http.NewRequest(http.MethodPost, revokeURL, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("client_id", t.Client.ID)
	q.Add("client_secret", t.Client.Secret)
	q.Add("token", t.User.Bearer)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", "Bearer "+t.User.Bearer)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to revoke token (status: %d)", res.StatusCode)
	}

	t.User.Bearer = ""
	return nil
}
