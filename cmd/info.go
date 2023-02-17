package cmd

import (
	"os"

	"github.com/arrow2nd/anct/cmdutil"
	"github.com/arrow2nd/anct/view"
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
	annictID, _, err := cmdutil.SearchWorks(c.api, cmd, args)
	if err != nil {
		return err
	}

	info, err := c.api.FetchWorkInfo(annictID)

	return view.PrintWorkInfo(os.Stdout, info)
}
