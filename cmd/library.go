package cmd

import (
	"context"
	"os"

	"github.com/arrow2nd/anct/api"
	"github.com/arrow2nd/anct/cmdutil"
	"github.com/arrow2nd/anct/gen"
	"github.com/arrow2nd/anct/view"
	"github.com/spf13/cobra"
)

func (c *Command) newCmdLibrary() *cobra.Command {
	library := &cobra.Command{
		Use:   "library",
		Short: "View own library",
		RunE:  c.libraryRun,
	}

	library.Flags().StringP("status", "s", "", "Library status state: {wanna_watch|watching|watched|on_hold|stop_watching}")
	library.Flags().StringP("from", "", "", "Retrieve works from a given season: YYYY-{spring|summer|autumn|winter}")
	library.Flags().StringP("until", "", "", "Retrieve works until a given season: YYYY-{spring|summer|autumn|winter}")
	cmdutil.SetLimitFlag(library.Flags())

	return library
}

func (c *Command) libraryRun(cmd *cobra.Command, arg []string) error {
	fromStr, _ := cmd.Flags().GetString("from")
	untilStr, _ := cmd.Flags().GetString("until")
	from := &fromStr
	until := &untilStr

	// シーズン指定の書式をチェック
	for _, s := range []**string{&from, &until} {
		if **s == "" {
			// NOTE: APIクライアントの仕様上、指定しない場合はnilを渡す必要がある
			*s = nil
			continue
		}

		if err := cmdutil.ValidateSeasonFormat(**s); err != nil {
			return err
		}
	}

	status, err := cmdutil.ReceiveStatus(cmd.Flags(), false)
	limit, _ := cmd.Flags().GetInt64("limit")

	list, err := c.api.Client.FetchUserLibrary(context.Background(), status, from, until, limit)
	if err != nil {
		return api.HandleClientError(err)
	}

	works := []*gen.WorkFragment{}
	for _, node := range list.Viewer.LibraryEntries.Nodes {
		if node != nil {
			works = append(works, node.Work)
		}
	}

	view.PrintWorksTable(os.Stdout, works)
	return nil
}
