package cmd

import (
	"github.com/arrow2nd/anct/cmdutil"
	"github.com/arrow2nd/anct/view"
	"github.com/spf13/cobra"
)

func (c *Command) newCmdStatus() *cobra.Command {
	s := &cobra.Command{
		Use:     "status [<query>]",
		Short:   "Update the watching status of work",
		Example: "  anct status ぼっち・ざ・ろっく！",
		RunE:    c.updateStatusRun,
	}

	cmdutil.SetSearchFlags(s.Flags())
	s.Flags().StringP("state", "s", "", "Update status state: {wanna_watch|watching|watched|on_hold|stop_watching|no_state}")

	return s
}

func (c *Command) updateStatusRun(cmd *cobra.Command, args []string) error {
	_, id, err := cmdutil.SearchWorks(c.api, cmd, args)
	if err != nil {
		return err
	}

	// 視聴状態を取得
	stateStr, _ := cmd.Flags().GetString("state")
	if stateStr == "" {
		// 指定されていなければ対話形式で聞く
		s, err := view.SelectStatus(true)
		if err != nil {
			return err
		}
		stateStr = s
	}

	state, err := cmdutil.StringToStatusState(stateStr, true)
	if err != nil {
		return err
	}

	spinner := view.SpinnerStart(cmd.OutOrStdout(), "Updating status")

	if err := c.api.UpdateWorkState(id, state); err != nil {
		return err
	}

	spinner.Stop()

	view.PrintDone(cmd.OutOrStdout(), "Updated status!")
	return nil
}
