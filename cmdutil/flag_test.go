package cmdutil

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestCommonFlags(t *testing.T) {
	p := pflag.NewFlagSet("test", pflag.ExitOnError)
	SetCommonFlags(p)

	editor, limit := GetCommonFlags(p)

	assert.False(t, editor, "editor = %v", editor)
	assert.Equal(t, int64(30), limit)
}

func TestAllSearchFlags(t *testing.T) {
	p := pflag.NewFlagSet("test", pflag.ExitOnError)
	SetSearchFlags(p)

	states, seasons, limit, useEditor, err := getAllSearchFlags(p)
	assert.NoError(t, err)

	assert.Len(t, states, 0)
	assert.Len(t, seasons, 0)
	assert.Equal(t, int64(30), limit)
	assert.False(t, useEditor)
}
