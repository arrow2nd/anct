package cmdutil

import (
	"errors"

	"github.com/arrow2nd/anct/api"
	"github.com/arrow2nd/anct/gen"
	"github.com/arrow2nd/anct/view"
	"github.com/spf13/cobra"
)

// SearchWorks : 作品を検索してIDを取得
func SearchWorks(api *api.API, cmd *cobra.Command, args []string) (int64, string, error) {
	states, seasons, limit, useEditor, err := getAllSearchFlags(cmd.Flags())
	if err != nil {
		return 0, "", err
	}

	// 検索クエリ
	query, err := receiveQuery("Search query", args, useEditor, false)
	if err != nil {
		return 0, "", err
	}

	query = StripWhiteSpace(query)

	// 条件指定が無い場合はエラー
	if query == "" && len(states) == 0 && len(seasons) == 0 {
		return 0, "", errors.New("query or `--library` or `--seasons` is required")
	}

	spinner := view.SpinnerStart(cmd.OutOrStdout(), "Searching for works")
	list := []*gen.WorkFragment{}
	err = nil

	// 作品を検索
	if len(states) == 0 {
		list, err = api.SearchWorks(query, seasons, limit)
	} else {
		list, err = api.SearchWorksFromLibrary(query, states, seasons, limit)
	}

	spinner.Stop()

	if err != nil {
		return 0, "", err
	}

	return view.SelectWork(list)
}
