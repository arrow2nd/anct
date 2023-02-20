package api

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	builtInClientID     = ""
	builtInClientSecret = ""
)

// ClientToken : クライアントトークン
type ClientToken struct {
	ID     string
	Secret string
}

func (ct *ClientToken) Get() (string, string, error) {
	if ct.InEmpty() {
		return "", "", errors.New("Client token not set. please run `anct config client-token` to set the token")
	}

	return ct.ID, ct.Secret, nil
}

func (ct *ClientToken) Set(id, secret string) {
	ct.ID = id
	ct.Secret = secret
}

// InEmpty : 値が設定されているか
func (ct *ClientToken) InEmpty() bool {
	return ct.ID == "" || ct.Secret == ""
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

	id, secret, err := t.Client.Get()
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("client_id", id)
	q.Add("client_secret", secret)
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
