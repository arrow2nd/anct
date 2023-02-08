package cmd

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/arrow2nd/anct/view"
	"github.com/spf13/cobra"
)

func (c *Command) newCmdSearchWorks() *cobra.Command {
	works := &cobra.Command{
		Use:     "works [<keyword>]",
		Short:   "Search for works",
		Example: "  anct search works ARIA --seasons 2005-autumn",
		RunE:    c.searchWorksRun,
	}

	setLimitFlag(works.Flags())
	setEditerFlag(works.Flags())
	works.Flags().StringSliceP("seasons", "s", []string{}, "Broadcast seasons: YYYY-{spring|summer|autumn|winter}")

	return works
}

func (c *Command) searchWorksRun(cmd *cobra.Command, args []string) error {
	useEditor, _ := cmd.Flags().GetBool("editor")
	limit, _ := cmd.Flags().GetInt64("limit")
	seasons, _ := cmd.Flags().GetStringSlice("seasons")

	// シーズン指定の書式をチェック
	for _, s := range seasons {
		if err := checkSeasonFormat(s); err != nil {
			return err
		}
	}

	keyword, err := receivekeyword(args, useEditor, true)
	if err != nil {
		return err
	}
	if keyword == "" && len(seasons) == 0 {
		return errors.New("keyword or `--seasons` is required")
	}

	ctx := context.Background()
	keywords := strings.Split(keyword, " ")
	list, err := c.api.Client.SearchWorksByKeyword(ctx, keywords, seasons, limit)
	if err != nil {
		// TODO: code = 401 のときに You are not logged in. Please run `anct auth login` を返したい
		return err
	}

	view.PrintWorksTable(os.Stdout, list.SearchWorks.Nodes)
	return nil
}