package view

import (
	"testing"

	"github.com/arrow2nd/anct/gen"
	"github.com/stretchr/testify/assert"
)

func TestCreateEpisodeOpt(t *testing.T) {
	mock := func() *gen.EpisodeFragment {
		n := int64(1)
		nt := "#1"
		t := "title"

		return &gen.EpisodeFragment{
			ID:                 "id",
			Number:             &n,
			NumberText:         &nt,
			Title:              &t,
			ViewerRecordsCount: 0,
		}
	}

	t.Run("作成できるか", func(t *testing.T) {
		opt := createEpisodeOpt(mock(), false)
		assert.Equal(t, "#1 title", opt, "話数とタイトル")
	})

	t.Run("記録件数も表示できるか", func(t *testing.T) {
		m := mock()
		m.ViewerRecordsCount = 1
		opt := createEpisodeOpt(m, true)
		assert.Equal(t, "#1 title - Recorded (1)", opt)
	})

	t.Run("話数が不明", func(t *testing.T) {
		m := mock()
		m.NumberText = nil
		opt := createEpisodeOpt(m, false)
		assert.Equal(t, "??? title", opt)
	})

	t.Run("タイトルが不明", func(t *testing.T) {
		m := mock()
		m.Title = nil
		opt := createEpisodeOpt(m, false)
		assert.Equal(t, "#1 ??? (ID: id)", opt)
	})
}
