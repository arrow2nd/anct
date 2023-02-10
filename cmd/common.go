package cmd

import (
	"errors"

	"github.com/arrow2nd/anct/cmdutil"
	"github.com/arrow2nd/anct/gen"
	"github.com/arrow2nd/anct/view"
	"github.com/spf13/cobra"
)

// searchWorks : 作品を検索してIDを取得
func (c *Command) searchWorks(cmd *cobra.Command, args []string) (string, error) {
	// 検索関連フラグの内容を取得
	useEditor, limit := cmdutil.ReceiveCommonFlags(cmd.Flags())
	states, seasons, err := cmdutil.ReceiveSearchCommonFlags(cmd.Flags())
	if err != nil {
		return "", err
	}

	// 検索クエリを受取る
	query, err := cmdutil.ReceiveQuery(args, useEditor, true)
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
		list, err = c.api.SearchWorks(query, seasons, limit)
	} else {
		list, err = c.api.SearchWorksFromLibrary(query, states, seasons, limit)
	}

	if err != nil {
		return "", err
	}

	// fzf で絞り込む
	return view.SelectWork(list)
}
