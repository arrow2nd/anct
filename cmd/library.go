package cmd

import "github.com/spf13/cobra"

func (a *App) newCmdLibrary() *cobra.Command {
	library := &cobra.Command{
		Use:   "library",
		Short: "View own library",
	}

	watching := &cobra.Command{
		Use:   "watching",
		Short: "Watching",
	}

	wannnaWatch := &cobra.Command{
		Use:   "wannna-watch",
		Short: "Wanna watch",
	}

	watched := &cobra.Command{
		Use:   "watched",
		Short: "Watched",
	}

	onHold := &cobra.Command{
		Use:   "on-hold",
		Short: "On hold",
	}

	stopWatching := &cobra.Command{
		Use:   "stop-watching",
		Short: "Stop watching",
	}

	library.AddCommand(
		watching,
		wannnaWatch,
		watched,
		onHold,
		stopWatching,
	)

	return library
}
