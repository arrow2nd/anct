package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (c *Command) newCmdSearch() *cobra.Command {
	setFlags := func(p *pflag.FlagSet) {
		p.BoolP("editor", "e", false, "Use an external editor to enter keyword")
		setCommonFlags(p)
	}

	works := c.newCmdSearchWorks()
	setFlags(works.Flags())

	characters := c.newCmdSearchCharacters()
	setFlags(characters.Flags())

	search := &cobra.Command{
		Use:   "search",
		Short: "Search for works, characters",
	}
	search.AddCommand(works, characters)

	return search
}
