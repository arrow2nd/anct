package cmd

import (
	"github.com/arrow2nd/anct/client"
	"github.com/spf13/cobra"
)

// App : 本体
type App struct {
	root   *cobra.Command
	client *client.Client
}

// New : 作成
func New(t *client.Token) *App {
	a := &App{
		root: &cobra.Command{
			Use:          "anct",
			Short:        "🎦 Unofficial CLI Client of Annict",
			SilenceUsage: true,
		},
		client: client.New(t),
	}

	a.root.AddCommand(
		a.newAuthCmd(),
		a.newSearchCmd(),
		a.newLibraryCmd(),
		a.newRecordCmd(),
		a.newVersionCmd(),
	)

	return a
}

// Execute : 実行
func (c *App) Execute() error {
	return c.root.Execute()
}
