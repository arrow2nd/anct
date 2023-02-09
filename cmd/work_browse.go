package cmd

import (
	"fmt"
	"io"

	"github.com/arrow2nd/anct/cmdutil"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

func (c *Command) newCmdWorkBrowse() *cobra.Command {
	browse := &cobra.Command{
		Use:   "browse <work id>",
		Short: "Open Annict's work page in the web browser",
		Args:  cobra.ExactArgs(1),
		RunE:  c.workBrowseRun,
	}

	return browse
}

func (c *Command) workBrowseRun(cmd *cobra.Command, args []string) error {
	workID, err := cmdutil.StringToWorkID(args[0])
	if err != nil {
		return err
	}

	browser.Stdout = io.Discard
	browser.Stderr = io.Discard

	url := fmt.Sprintf("https://annict.com/works/%d", workID)
	browser.OpenURL(url)

	return nil
}
