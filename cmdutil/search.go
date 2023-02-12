package cmdutil

import (
	"errors"

	"github.com/arrow2nd/anct/api"
	"github.com/arrow2nd/anct/gen"
	"github.com/arrow2nd/anct/view"
	"github.com/spf13/cobra"
)

// SearchWorks : 作品を検索してIDを取得
func SearchWorks(api *api.API, cmd *cobra.Command, args []string) (string, error) {
	// 検索関連フラグの内容を取得
	states, seasons, limit, useEditor, err := ReceiveSearchFlags(cmd.Flags())
	if err != nil {
		return "", err
	}

	// 検索クエリを受取る
	query, err := ReceiveQuery(args, useEditor, true)
	if err != nil {
		return "", err
	}

	// 条件指定が無い場合はエラー
	if query == "" && len(states) == 0 && len(seasons) == 0 {
		return "", errors.New("query or `--library` or `--seasons` is required")
	}

	// 作品を検索
	list := []*gen.WorkFragment{}
	err = nil

	if len(states) == 0 {
		list, err = api.SearchWorks(query, seasons, limit)
	} else {
		list, err = api.SearchWorksFromLibrary(query, states, seasons, limit)
	}

	if err != nil {
		return "", err
	}

	return view.SelectWork(list)
}
