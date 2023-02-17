package cmd

import (
	"github.com/arrow2nd/anct/cmdutil"
	"github.com/spf13/cobra"
)

func (c *Command) newCmdRecord() *cobra.Command {
	record := &cobra.Command{
		Use:     "record [<query>]",
		Short:   "Record the watching of episode",
		Example: "  anct record お兄ちゃんはおしまい",
		RunE:    c.recordRun,
	}

	return record
}

func (c *Command) recordRun(cmd *cobra.Command, args []string) error {
	_, id, err := cmdutil.SearchWorks(c.api, cmd, args)
	if err != nil {
		return err
	}

	return nil
}
