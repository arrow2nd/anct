package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "develop"

func (c *Cmd) newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Aliases: []string{"ver"},
		Short:   "Display current version",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("🎦 anct ver.%s\n", version)
		},
	}
}
