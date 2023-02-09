package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (c *Command) newCmdSearchCharacters() *cobra.Command {
	characters := &cobra.Command{
		Use:     "characters [<keyword>]",
		Short:   "Search for characters",
		Example: "  anct search characters 後藤ひとり",
		RunE:    c.searchCharactersRun,
	}

	setLimitFlag(characters.Flags())
	setEditerFlag(characters.Flags())

	return characters
}

func (c *Command) searchCharactersRun(cmd *cobra.Command, args []string) error {
	// NOTE: 以下の issue のため、このコマンドは無効化しています
	//       https://github.com/arrow2nd/anct/issues/2

	fmt.Println(`This command cannot be used because of the following issue.
https://github.com/arrow2nd/anct/issues/2`)

	// useEditor, _ := cmd.Flags().GetBool("editor")
	// limit, _ := cmd.Flags().GetInt64("limit")

	// keyword, err := view.Receivekeyword(args, useEditor, false)
	// if err != nil {
	// 	return err
	// }

	// ctx := context.Background()
	// list, err := c.api.Client.SearchCharactersByKeyword(ctx, keyword, limit)
	// if err != nil {
	// return api.HandleClientError(err)
	// }

	// view.PrintCharactersTable(os.Stdout, list.SearchCharacters.Nodes)
	return nil
}
