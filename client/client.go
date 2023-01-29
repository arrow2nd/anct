package client

// Client : クライアント
type Client struct {
	Token Token
}

func New(t *Token) *Client {
	c := &Client{
		Token: *t,
	}

	if c.Token.Client.ID == "" || c.Token.Client.Secret == "" {
		c.Token.Client.ID = clientID
		c.Token.Client.Secret = clientSecret
	}

	return c
}
