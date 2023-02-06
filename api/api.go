package api

import (
	"context"
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/arrow2nd/anct/gen"
)

const baseURL = "https://api.annict.com/graphql"

// API : APIクライアント
type API struct {
	Client *gen.Client
	Token  Token
}

func New(t *Token) *API {
	ac := gen.NewClient(http.DefaultClient, baseURL, func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {
		req.Header.Set("Authorization", "Bearer "+t.User.Bearer)
		return next(ctx, req, gqlInfo, res)
	})

	c := &API{
		Client: ac,
		Token:  *t,
	}

	// 組込みのクライアントトークンを使う
	if c.Token.Client.ID == "" || c.Token.Client.Secret == "" {
		c.Token.Client.ID = clientID
		c.Token.Client.Secret = clientSecret
	}

	return c
}
