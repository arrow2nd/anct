package cmd

import (
	"fmt"

	"github.com/arrow2nd/anct/cmdutil"
	"github.com/spf13/cobra"
)

func (c *Command) newCmdSearch() *cobra.Command {
	works := &cobra.Command{
		Use:     "search [<query>]",
		Short:   "Search for works",
		Example: "  anct search ARIA --seasons 2005-autumn",
		RunE:    c.searchWorksRun,
	}

	cmdutil.SetSearchFlags(works.Flags())

	return works
}

func (c *Command) searchWorksRun(cmd *cobra.Command, args []string) error {
	id, err := c.searchWorks(cmd, args)
	if err != nil {
		return err
	}

	// TODO: 要らないかも

	fmt.Println("GET: " + id)
	return nil
}
