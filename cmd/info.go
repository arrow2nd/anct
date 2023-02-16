package cmd

import (
	"github.com/arrow2nd/anct/cmdutil"
	"github.com/spf13/cobra"
)

func (c *Command) newCmdInfo() *cobra.Command {
	info := &cobra.Command{
		Use:     "info [<query>]",
		Short:   "View information about the work",
		Example: "  anct info ARIA --seasons 2005-autumn",
		RunE:    c.infoRun,
	}

	cmdutil.SetSearchFlags(info.Flags())

	return info
}

func (c *Command) infoRun(cmd *cobra.Command, args []string) error {
	annictID, err := cmdutil.SearchWorks(c.api, cmd, args)
	if err != nil {
		return err
	}

	info, err := c.api.FetchWorkInfo(annictID)

	return nil
}
