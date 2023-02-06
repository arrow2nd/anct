package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func (c *Command) newCmdSearch() *cobra.Command {
	search := &cobra.Command{
		Use:   "search",
		Short: "Search for works, characters",
	}

	works := &cobra.Command{
		Use:   "works [<query>]",
		Short: "Search for works",
		Args:  cobra.ExactArgs(1),
		RunE:  c.searchWorksRun,
	}

	characters := &cobra.Command{
		Use:   "characters [<query>]",
		Short: "Search fot characters",
		Args:  cobra.ExactArgs(1),
	}

	search.AddCommand(
		works,
		characters,
	)

	return search
}

func (c *Command) searchWorksRun(cmd *cobra.Command, arg []string) error {
	ctx := context.Background()
	list, err := c.api.Client.SearchWorksByKeyword(ctx, arg[0], 5)
	if err != nil {
		// TODO: code = 401 のときに You are not logged in. Please run `anct auth login` を返したい
		return err
	}

	for _, node := range list.SearchWorks.Nodes {
		fmt.Printf("Work ID: %d / TITLE: %s / MEDIA: %s / SEASON: %d %s\n", node.AnnictID, node.Title, node.Media.String(), *node.SeasonYear, node.SeasonName.String())
	}

	return nil
}
