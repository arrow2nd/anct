package cmd

import (
	"github.com/arrow2nd/anct/cmdutil"
	"github.com/arrow2nd/anct/view"
	"github.com/spf13/cobra"
)

func (c *Command) newCmdState() *cobra.Command {
	u := &cobra.Command{
		Use:  "status [<query>]",
		RunE: c.updateState,
	}

	u.Flags().StringP("state", "", "", "Update status state: {wanna_watch|watching|watched|on_hold|stop_watching|no_state}")
	cmdutil.SetSearchFlags(u.Flags())

	return u
}

func (c *Command) updateState(cmd *cobra.Command, args []string) error {
	id, err := cmdutil.SearchWorks(c.api, cmd, args)
	if err != nil {
		return err
	}

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

	return c.api.UpdateWorkState(id, state)
}
