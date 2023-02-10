package api

import (
	"fmt"

	"github.com/Yamashou/gqlgenc/clientv2"
)

// handleClientError : Clientのネットワークエラーをハンドリング
func handleClientError(e error) error {
	if hErr, ok := e.(*clientv2.ErrorResponse); ok {
		return fmt.Errorf("%s (status code: %d)", hErr.NetworkError.Message, hErr.NetworkError.Code)
	}
	return e
}
