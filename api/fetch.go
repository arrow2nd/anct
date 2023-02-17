package api

import (
	"context"
	"fmt"

	"github.com/arrow2nd/anct/gen"
)

// FetchWorkInfo : 作品の詳細を取得
func (a *API) FetchWorkInfo(annictID int64) (*gen.WorkInfoFragment, error) {
	ctx := context.Background()

	result, err := a.client.FetchWorkInfo(ctx, annictID)
	if err != nil {
		return nil, err
	}

	if len(result.SearchWorks.Nodes) == 0 {
		return nil, fmt.Errorf("not found work (annictID: %d)", annictID)
	}

	return result.SearchWorks.Nodes[0], nil
}

// FetchWorkEpisodes : 作品のエピソードを取得
func (a *API) FetchWorkEpisodes(annictID int64) (*gen.WorkEpisodesFragment, error) {
	ctx := context.Background()

	result, err := a.client.FetchWorkEpisodes(ctx, annictID)
	if err != nil {
		return nil, err
	}

	if len(result.SearchWorks.Nodes) == 0 {
		return nil, fmt.Errorf("not found work (annictID: %d)", annictID)
	}

	return result.SearchWorks.Nodes[0], nil
}
