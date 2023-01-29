package cmd

import "github.com/spf13/cobra"

// Cmd : 本体
type Cmd struct {
	root *cobra.Command
}

// New : 作成
func New() *Cmd {
	c := &Cmd{
		root: &cobra.Command{
			Use:          "anct",
			Short:        "🎦 Unofficial CLI Client of Annict",
			SilenceUsage: true,
		},
	}

	c.root.AddCommand(
		c.newAuthCmd(),
		c.newSearchCmd(),
		c.newLibraryCmd(),
		c.newRecordCmd(),
		c.newVersionCmd(),
	)

	return c
}

// Execute : 実行
func (c *Cmd) Execute() error {
	return c.root.Execute()
}
