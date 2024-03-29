package cmd

import (
	"github.com/arrow2nd/anct/cmdutil"
	"github.com/arrow2nd/anct/view"
	"github.com/spf13/cobra"
)

func (c *Command) newCmdInfo() *cobra.Command {
	info := &cobra.Command{
		Use:     "info [<query>]",
		Short:   "Display information about the work",
		Example: "  anct info ARIA --seasons 2005-autumn",
		RunE:    c.infoRun,
	}

	cmdutil.SetSearchFlags(info.Flags())

	return info
}

func (c *Command) infoRun(cmd *cobra.Command, args []string) error {
	work, _, err := cmdutil.SearchWorks(c.api, cmd, args)
	if err != nil {
		return err
	}

	spinner := view.SpinnerStart(cmd.OutOrStdout(), "Loading information the work")
	info, err := c.api.FetchWorkInfo(work.AnnictID)
	if err != nil {
		return err
	}

	spinner.Stop()
	return view.PrintWorkInfo(cmd.OutOrStdout(), info)
}
