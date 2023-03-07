package cmdutil

import (
	"testing"

	"github.com/arrow2nd/anct/gen"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestReceiveQuery(t *testing.T) {
	t.Run("引数から受け取る", func(t *testing.T) {
		args := []string{"a", "b", "c"}
		q, err := receiveQuery("test", args, false)

		assert.NoError(t, err)
		assert.Equal(t, "a b c", q)
	})
}

func TestReceiveRating(t *testing.T) {
	t.Run("フラグから受け取る", func(t *testing.T) {
		p := pflag.NewFlagSet("test", pflag.ExitOnError)
		p.StringP("rating", "r", "", "episode rating: {great|good|average|bad}")

		err := p.Parse([]string{"--rating", "great"})
		assert.NoError(t, err)

		rating, err := ReceiveRating(p, "rating")
		assert.NoError(t, err)
		assert.Equal(t, gen.RatingStateGreat, rating)
	})
}

func TestReceiveBody(t *testing.T) {
	t.Run("フラグから受け取る", func(t *testing.T) {
		p := pflag.NewFlagSet("test", pflag.ExitOnError)
		p.StringP("body", "b", "", "test")

		want := "body"
		err := p.Parse([]string{"--body", want})
		assert.NoError(t, err)

		got, err := ReceiveBody(p, "body")
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})
}
