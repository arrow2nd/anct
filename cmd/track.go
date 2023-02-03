package cmd

import "github.com/spf13/cobra"

func (a *App) newCmdRecord() *cobra.Command {
	return &cobra.Command{
		Use:   "record",
		Short: "Record unwatched episodes",
		Args:  cobra.NoArgs,
	}
}
