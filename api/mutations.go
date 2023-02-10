package api

import (
	"context"

	"github.com/arrow2nd/anct/gen"
)

// UpdateWorkState : 作品の視聴ステータスを更新
func (a *API) UpdateWorkState(id string, state gen.StatusState) error {
	ctx := context.Background()
	if _, err := a.client.UpdateWorkState(ctx, id, state); err != nil {
		return err
	}

	return nil
}
