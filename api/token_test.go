package api_test

import (
	"testing"

	"github.com/arrow2nd/anct/api"
	"github.com/stretchr/testify/assert"
)

func TestClientToken(t *testing.T) {
	t.Run("クライアントトークンがない場合エラーを返す", func(t *testing.T) {
		ct := &api.ClientToken{}
		_, _, err := ct.Get()
		assert.Error(t, err)
	})

	t.Run("設定した値が取得できるか", func(t *testing.T) {
		ct := &api.ClientToken{}
		ct.Set("id", "secret")

		id, secret, err := ct.Get()
		assert.NoError(t, err)
		assert.Equal(t, "id", id)
		assert.Equal(t, "secret", secret)
	})
}

func TestClientTokenIsEmpty(t *testing.T) {
	ct := &api.ClientToken{}
	assert.True(t, ct.InEmpty(), "空")

	ct.Set("a", "b")
	assert.False(t, ct.InEmpty(), "空ではない")
}
