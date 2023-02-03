package cmd

import "github.com/spf13/cobra"

func (c *Command) newCmdSearch() *cobra.Command {
	search := &cobra.Command{
		Use:   "search",
		Short: "Search for works, characters",
	}

	characters := &cobra.Command{
		Use:   "characters [<query>]",
		Short: "Search fot characters",
		Args:  cobra.ExactArgs(1),
	}

	works := &cobra.Command{
		Use:   "works [<query>]",
		Short: "Search for works",
		Args:  cobra.ExactArgs(1),
	}

	search.AddCommand(
		characters,
		works,
	)

	return search
}
