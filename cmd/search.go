package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/arrow2nd/anct/view"
	"github.com/spf13/cobra"
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
		RunE:  c.searchCharactersRun,
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

	view.PrintWorksTable(os.Stdout, list.SearchWorks.Nodes)
	return nil
}

func (c *Command) searchCharactersRun(cmd *cobra.Command, arg []string) error {

	// NOTE: 以下の issue のため、このコマンドは無効化しています
	fmt.Println(`This command cannot be used because of the following issue.
https://github.com/arrow2nd/anct/issues/2`)

	// keyword, err := receivekeyword(arg)
	// if err != nil {
	// 	return err
	// }

	// limit, _ := cmd.Flags().GetInt64("limit")
	// ctx := context.Background()

	// list, err := c.api.Client.SearchCharactersByKeyword(ctx, keyword, limit)
	// if err != nil {
	// 	return err
	// }

	// view.PrintCharactersTable(os.Stdout, list.SearchCharacters.Nodes)
	return nil
}
