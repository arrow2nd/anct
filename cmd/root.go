package cmd

import (
	"github.com/arrow2nd/anct/api"
	"github.com/arrow2nd/anct/config"
	"github.com/spf13/cobra"
)

// Command : コマンド本体
type Command struct {
	root *cobra.Command
	api  *api.API
	cfg  *config.Config
}

// New : 作成
func New(cfg *config.Config) (*Command, error) {
	t, err := cfg.Load()
	if err != nil {
		return nil, err
	}

	c := &Command{
		root: &cobra.Command{
			Use:           "anct",
			Short:         "📺 Unofficial CLI Client of Annict",
			SilenceUsage:  true,
			SilenceErrors: true,
		},
		api: api.New(t),
		cfg: cfg,
	}

	c.root.AddCommand(
		c.newCmdAuth(),
		c.newCmdConfig(),
		c.newCmdInfo(),
		c.newCmdStatus(),
		c.newCmdRecord(),
		c.newCmdReview(),
		c.newCmdVersion(),
	)

	return c, nil
}

// Execute : 実行
func (c *Command) Execute() error {
	return c.root.Execute()
}
