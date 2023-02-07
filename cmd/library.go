package cmd

import (
	"context"
	"os"

	"github.com/arrow2nd/anct/gen"
	"github.com/arrow2nd/anct/view"
	"github.com/spf13/cobra"
)

func (c *Command) newCmdLibrary() *cobra.Command {
	library := &cobra.Command{
		Use:   "library [watch status]",
		Short: "View own library",
		RunE:  c.libraryRun,
	}

	return library
}

func (c *Command) libraryRun(cmd *cobra.Command, arg []string) error {
	status, err := view.SelectWatchState()
	if err != nil {
		return err
	}

	ctx := context.Background()
	list, err := c.api.Client.FetchUserLibrary(ctx, status, 30)
	if err != nil {
		return err
	}

	works := []*gen.WorkFragment{}
	for _, node := range list.Viewer.LibraryEntries.Nodes {
		if node != nil {
			works = append(works, node.Work)
		}
	}

	view.PrintWorksTable(os.Stdout, works)
	return nil
}
