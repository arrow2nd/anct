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
		Use:   "library",
		Short: "View own library",
		RunE:  c.libraryRun,
	}

	library.Flags().StringP("status", "s", "", "Status state: {wanna_watch|watching|watched|on_hold|stop_watching}")
	setLimitFlag(library.Flags())

	return library
}

func (c *Command) libraryRun(cmd *cobra.Command, arg []string) error {
	limit, _ := cmd.Flags().GetInt64("limit")
	statusStr, _ := cmd.Flags().GetString("status")

	// フラグで指定されていない場合、対話形式で聞く
	if statusStr == "" {
		s, err := view.SelectStatus(false)
		if err != nil {
			return err
		}
		statusStr = s
	}

	status, err := toStatusState(statusStr, false)
	if err != nil {
		return err
	}

	ctx := context.Background()
	list, err := c.api.Client.FetchUserLibrary(ctx, status, limit)
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
