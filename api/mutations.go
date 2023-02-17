package api

import (
	"context"

	"github.com/arrow2nd/anct/gen"
	"golang.org/x/sync/errgroup"
)

// UpdateWorkState : 作品の視聴ステータスを更新
func (a *API) UpdateWorkState(id string, state gen.StatusState) error {
	ctx := context.Background()
	if _, err := a.client.UpdateWorkState(ctx, id, state); err != nil {
		return err
	}

	return nil
}

// CreateEpisodeRecords : エピソードの視聴記録を作成
func (a *API) CreateEpisodeRecords(episodeIDs []string, rating gen.RatingState, comment string) error {
	eg, ctx := errgroup.WithContext(context.Background())

	for _, ID := range episodeIDs {
		ID := ID

		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				if _, err := a.client.CreateEpisodeRecord(context.Background(), ID, rating, &comment); err != nil {
					return err
				}
				return nil
			}
		})
	}

	return eg.Wait()
}
