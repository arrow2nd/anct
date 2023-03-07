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
	client *gen.Client
	Token  Token
}

// New : 新しいクライアントを作成
func New(t *Token) *API {
	header := func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {
		req.Header.Set("Authorization", "Bearer "+t.User.Bearer)
		return next(ctx, req, gqlInfo, res)
	}

	ac := gen.NewClient(http.DefaultClient, baseURL, nil, header)
	c := &API{
		client: ac,
		Token:  *t,
	}

	// 組込みのクライアントトークンを設定
	if c.Token.Client.InEmpty() {
		c.Token.Client.Set(builtInClientID, builtInClientSecret)
	}

	return c
}
