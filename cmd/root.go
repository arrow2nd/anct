package cmd

import (
	"github.com/arrow2nd/anct/api"
	"github.com/spf13/cobra"
)

// Command : コマンド本体
type Command struct {
	root   *cobra.Command
	client *api.Client
}

// New : 作成
func New(t *api.Token) *Command {
	c := &Command{
		root: &cobra.Command{
			Use:           "anct",
			Short:         "🎦 Unofficial CLI Client of Annict",
			SilenceUsage:  true,
			SilenceErrors: true,
		},
		client: api.NewClient(t),
	}

	c.root.AddCommand(
		c.newCmdAuth(),
		c.newCmdSearch(),
		c.newCmdLibrary(),
		c.newCmdVersion(),
	)

	return c
}

// Execute : 実行
func (c *Command) Execute() error {
	return c.root.Execute()
}
