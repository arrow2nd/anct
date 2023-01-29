package cmd

import "github.com/spf13/cobra"

func (c *App) newSearchCmd() *cobra.Command {
	search := &cobra.Command{
		Use:   "search",
		Short: "Search for animes, characters",
	}

	characters := &cobra.Command{
		Use:   "characters",
		Short: "Search fot characters",
		Args:  cobra.ExactValidArgs(1),
	}

	works := &cobra.Command{
		Use:   "works",
		Short: "Search for works",
		Args:  cobra.ExactArgs(1),
	}

	search.AddCommand(
		characters,
		works,
	)

	return search
}
