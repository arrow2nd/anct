package api

import "net/http"

// Credencial : 認証情報
type Credencial struct {
	// AccessToken : トークン
	AccessToken string `json:"access_token"`
}

// Revoke : トークンを失効させる
func (c *Credencial) Revoke() error {
	req, err := http.NewRequest(http.MethodPost, revokeURL, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("client_id", cliendID)
	q.Add("client_secret", clientSecret)
	q.Add("token", c.AccessToken)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", "Bearer "+c.AccessToken)

	client := http.Client{}
	if _, err := client.Do(req); err != nil {
		return err
	}

	c.AccessToken = ""
	return nil
}
