package view_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/arrow2nd/anct/gen"
	"github.com/arrow2nd/anct/view"
	"github.com/stretchr/testify/assert"
)

func TestPrintWorkInfo(t *testing.T) {
	mock := func() *gen.WorkInfoFragment {
		seasonName := gen.SeasonNameSpring
		seasonYear := int64(2005)
		statusState := gen.StatusStateWatching

		epNumber := int64(1)
		epNumberText := "1"
		epTitle := "ep_title"
		officialSite := "official_url"
		return &gen.WorkInfoFragment{
			AnnictID:          12345,
			ID:                "id",
			Title:             "title",
			Media:             gen.MediaTv,
			SeasonName:        &seasonName,
			SeasonYear:        &seasonYear,
			ViewerStatusState: &statusState,
			Episodes: &gen.WorkInfoFragment_WorkEpisodesFragment_Episodes{
				Nodes: []*gen.EpisodeFragment{
					{
						ID:                 "id",
						Number:             &epNumber,
						NumberText:         &epNumberText,
						Title:              &epTitle,
						ViewerRecordsCount: 2,
					},
				},
			},
			NoEpisodes:      false,
			Image:           nil,
			OfficialSiteURL: &officialSite,
			WatchersCount:   5,
		}
	}

	t.Run("全ての情報が出力されているか", func(t *testing.T) {
		buf := &bytes.Buffer{}
		view.PrintWorkInfo(buf, mock())

		want := `DETAIL
------
   TITLE:  title
   MEDIA:  TV
   SEASON: 2005 SPRING
   URL:    official_url

DATA
----
   WORK ID:    id
   ANNICT ID:  12345
   ANNICT URL: https://annict.com/works/12345
   WATCHERS:   5
   STATUS:     WATCHING

EPISODES
--------
   1  ep_title`

		assert.Equal(t, want, strings.TrimSpace(buf.String()))
	})

	t.Run("シーズンとURLがない", func(t *testing.T) {
		info := mock()
		info.SeasonYear = nil
		info.SeasonName = nil
		info.OfficialSiteURL = nil

		buf := &bytes.Buffer{}
		view.PrintWorkInfo(buf, info)

		want := `DETAIL
------
   TITLE:  title
   MEDIA:  TV
   SEASON: unknown
   URL:    unknown

DATA
----
   WORK ID:    id
   ANNICT ID:  12345
   ANNICT URL: https://annict.com/works/12345
   WATCHERS:   5
   STATUS:     WATCHING

EPISODES
--------
   1  ep_title`

		assert.Equal(t, want, strings.TrimSpace(buf.String()))
	})

	t.Run("エピソードがない", func(t *testing.T) {
		info := mock()
		info.Episodes = &gen.WorkInfoFragment_WorkEpisodesFragment_Episodes{
			Nodes: []*gen.EpisodeFragment{},
		}

		buf := &bytes.Buffer{}
		view.PrintWorkInfo(buf, info)

		want := `DETAIL
------
   TITLE:  title
   MEDIA:  TV
   SEASON: 2005 SPRING
   URL:    official_url

DATA
----
   WORK ID:    id
   ANNICT ID:  12345
   ANNICT URL: https://annict.com/works/12345
   WATCHERS:   5
   STATUS:     WATCHING

EPISODES
--------
   None yet.`

		assert.Equal(t, want, strings.TrimSpace(buf.String()))
	})
}
