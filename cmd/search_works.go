package cmd

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/arrow2nd/anct/api"
	"github.com/arrow2nd/anct/cmdutil"
	"github.com/arrow2nd/anct/view"
	"github.com/spf13/cobra"
)

func (c *Command) newCmdSearchWorks() *cobra.Command {
	works := &cobra.Command{
		Use:     "works [<query>]",
		Short:   "Search for works",
		Example: "  anct search works ARIA --seasons 2005-autumn",
		RunE:    c.searchWorksRun,
	}

	cmdutil.SetSearchFlags(works.Flags())

	return works
}

func (c *Command) searchWorksRun(cmd *cobra.Command, args []string) error {
	useEditor, _ := cmd.Flags().GetBool("editor")
	limit, _ := cmd.Flags().GetInt64("limit")
	seasons, _ := cmd.Flags().GetStringSlice("seasons")

	// シーズン指定の書式をチェック
	for _, s := range seasons {
		if err := cmdutil.ValidateSeasonFormat(s); err != nil {
			return err
		}
	}

	keyword, err := cmdutil.Receivekeyword(args, useEditor, true)
	if err != nil {
		return err
	}
	if keyword == "" && len(seasons) == 0 {
		return errors.New("query or `--seasons` is required")
	}

	keywords := strings.Split(keyword, " ")
	list, err := c.api.Client.SearchWorksByKeyword(context.Background(), keywords, seasons, limit)
	if err != nil {
		return api.HandleClientError(err)
	}

	view.PrintWorksTable(os.Stdout, list.SearchWorks.Nodes)
	return nil
}
