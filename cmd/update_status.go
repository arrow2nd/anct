package cmd

import (
	"github.com/arrow2nd/anct/cmdutil"
	"github.com/spf13/cobra"
)

func (c *Command) newCmdUpdateStatus() *cobra.Command {
	u := &cobra.Command{
		Use:  "update-status [<query>]",
		RunE: c.updateStatusRun,
	}

	u.Flags().StringP("state", "s", "", "Update status state: {wanna_watch|watching|watched|on_hold|stop_watching|no_state}")
	u.Flags().StringSliceP("ids", "", []string{}, "Specify the work ID directly")
	cmdutil.SetSearchFlags(u.Flags())

	return u
}

func (c *Command) updateStatusRun(cmd *cobra.Command, args []string) error {
	// workID, err := cmdutil.StringToWorkID(args[0])
	// if err != nil {
	// 	return err
	// }

	// フラグで指定されていない場合、対話形式で聞く
	// s, err := view.SelectStatus(allowNoState)
	// if err != nil {
	// 	return "", err
	// }
	// stateStr = s

	// if _, err := c.api.Client.UpdateWorkState(context.Background(), workID, status); err != nil {
	// 	return api.HandleClientError(err)
	// }

	return nil
}
