package cmd

import "github.com/spf13/cobra"

func (c *Command) newCmdWork() *cobra.Command {
	work := &cobra.Command{
		Use:   "work",
		Short: "Perform operations related to the work",
	}

	work.AddCommand(
		c.newCmdWorkBrowse(),
	)

	return work
}
