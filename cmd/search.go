package cmd

import "github.com/spf13/cobra"

func (c *Command) newCmdSearch() *cobra.Command {
	search := &cobra.Command{
		Use:   "search",
		Short: "Search for works, characters",
	}

	search.AddCommand(
		c.newCmdSearchWorks(),
		c.newCmdSearchCharacters(),
	)

	return search
}
