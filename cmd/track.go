package cmd

import "github.com/spf13/cobra"

func (c *Cmd) newRecordCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "record",
		Short: "Record unwatched episodes",
		Args:  cobra.NoArgs,
	}
}
