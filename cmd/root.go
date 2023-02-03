package cmd

import (
	"github.com/arrow2nd/anct/api"
	"github.com/spf13/cobra"
)

// App : 本体
type App struct {
	root   *cobra.Command
	client *api.Client
}

// New : 作成
func New(t *api.Token) *App {
	a := &App{
		root: &cobra.Command{
			Use:           "anct",
			Short:         "🎦 Unofficial CLI Client of Annict",
			SilenceUsage:  true,
			SilenceErrors: true,
		},
		client: api.NewClient(t),
	}

	a.root.AddCommand(
		a.newCmdAuth(),
		a.newCmdSearch(),
		a.newCmdLibrary(),
		a.newCmdRecord(),
		a.newCmdVersion(),
	)

	return a
}

// Execute : 実行
func (a *App) Execute() error {
	return a.root.Execute()
}
