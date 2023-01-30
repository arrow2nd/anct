package api

// Client : クライアント
type Client struct {
	Token Token
}

func NewClient(t *Token) *Client {
	c := &Client{
		Token: *t,
	}

	if c.Token.Client.ID == "" || c.Token.Client.Secret == "" {
		c.Token.Client.ID = clientID
		c.Token.Client.Secret = clientSecret
	}

	return c
}
