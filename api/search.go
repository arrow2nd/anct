package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/arrow2nd/anct/gen"
)

// SearchWorks : クエリから作品を検索
func (a *API) SearchWorks(q string, seasons []string, limit int64) ([]*gen.WorkFragment, error) {
	ctx := context.Background()

	list, err := a.client.SearchWorksByKeyword(ctx, q, seasons, limit)
	if err != nil {
		return nil, handleClientError(err)
	}

	if len(list.SearchWorks.Nodes) == 0 {
		return nil, fmt.Errorf("no matching works found (query: %s)", q)
	}

	return list.SearchWorks.Nodes, nil
}

// SearchWorksFromLibrary : ライブラリ内の作品を検索
func (a *API) SearchWorksFromLibrary(q string, states []gen.StatusState, seasons []string, limit int64) ([]*gen.WorkFragment, error) {
	ctx := context.Background()

	list, err := a.client.FetchUserLibrary(ctx, states, seasons, limit)
	if err != nil {
		return nil, handleClientError(err)
	}

	// クエリで絞り込む
	works := []*gen.WorkFragment{}
	for _, node := range list.Viewer.LibraryEntries.Nodes {
		if node == nil {
			continue
		}

		if q == "" || strings.Contains(node.Work.Title, q) {
			works = append(works, node.Work)
		}
	}

	if len(works) == 0 {
		return nil, fmt.Errorf("no matches found in the library (states: %v)", states)
	}

	return works, nil
}
