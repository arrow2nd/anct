package cmdutil

import (
	"errors"

	"github.com/arrow2nd/anct/api"
	"github.com/arrow2nd/anct/gen"
	"github.com/arrow2nd/anct/view"
	"github.com/spf13/cobra"
	"github.com/usk81/r2h"
)

// SearchWorks : 作品を検索してIDを取得
func SearchWorks(api *api.API, cmd *cobra.Command, args []string) (*gen.WorkFragment, string, error) {
	states, seasons, limit, useEditor, err := getAllSearchFlags(cmd.Flags())
	if err != nil {
		return nil, "", err
	}

	// 検索クエリ
	query, err := receiveQuery("Search query", args, useEditor)
	if err != nil {
		return nil, "", err
	}

	query = StripWhiteSpace(query)

	// ローマ字 -> ひらがな変換
	convertResult, err := r2h.ConvertStrict(query)
	if err == nil {
		query = convertResult
	}

	// 条件指定が無い場合はエラー
	if query == "" && len(states) == 0 && len(seasons) == 0 {
		return nil, "", errors.New("query or `--library` or `--seasons` is required")
	}

	spinner := view.SpinnerStart(cmd.OutOrStdout(), "Searching for works")
	works := []*gen.WorkFragment{}
	err = nil

	// 作品を検索
	if len(states) == 0 {
		works, err = api.SearchWorks(query, seasons, limit)
	} else {
		works, err = api.SearchWorksFromLibrary(query, states, seasons, limit)
	}

	spinner.Stop()

	if err != nil {
		return nil, "", err
	}

	return view.SelectWork(works)
}
