package cmd

import (
	"context"
	"github.com/arrow2nd/anct/view"
	"github.com/spf13/cobra"
	"os"
)

func (c *Command) newCmdSearch() *cobra.Command {
	search := &cobra.Command{
		Use:   "search",
		Short: "Search for works, characters",
	}

	works := &cobra.Command{
		Use:   "works [keyword]",
		Short: "Search for works",
		RunE:  c.searchWorksRun,
	}
	setSearchFlags(works)

	characters := &cobra.Command{
		Use:   "characters [keyword]",
		Short: "Search for characters",
	}
	setSearchFlags(characters)

	search.AddCommand(
		works,
		characters,
	)

	return search
}

func setSearchFlags(cmd *cobra.Command) {
	cmd.Flags().Int64P("limit", "L", 30, "Maximum number of results to fetch")
}

func (c *Command) searchWorksRun(cmd *cobra.Command, arg []string) error {
	keyword, err := receivekeyword(arg)
	if err != nil {
		return err
	}

	limit, _ := cmd.Flags().GetInt64("limit")
	ctx := context.Background()

	list, err := c.api.Client.SearchWorksByKeyword(ctx, keyword, limit)
	if err != nil {
		// TODO: code = 401 のときに You are not logged in. Please run `anct auth login` を返したい
		return err
	}

	view.PrintWorksTable(os.Stdout, keyword, list.SearchWorks.Nodes)
	return nil
}
