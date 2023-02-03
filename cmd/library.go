package cmd

import "github.com/spf13/cobra"

func (c *Command) newCmdLibrary() *cobra.Command {
	library := &cobra.Command{
		Use:   "library [<watch status>]",
		Short: "View own library",
	}

	return library
}
