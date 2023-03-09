package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/arrow2nd/anct/api"
	"github.com/arrow2nd/anct/cmdutil"
	"github.com/arrow2nd/anct/view"
	"github.com/spf13/cobra"
)

func (c *Command) newCmdConfig() *cobra.Command {
	cf := &cobra.Command{
		Use:     "config",
		Short:   "Edit the configuration",
		Example: "  anct config",
	}

	ct := &cobra.Command{
		Use:     "client-token",
		Short:   "Change the API key used to connect to Annict",
		Example: "  anct config client-token",
		RunE:    c.conifgRun,
	}

	cf.AddCommand(ct)
	return cf
}

func (c *Command) conifgRun(cmd *cobra.Command, args []string) error {
	trans := func(ans interface{}) interface{} {
		if id, ok := ans.(string); ok {
			return cmdutil.StripWhiteSpace(id)
		}
		return ""
	}

	qs := []*survey.Question{
		{
			Name:      "id",
			Prompt:    &survey.Input{Message: "Client ID"},
			Transform: trans,
		},
		{
			Name:      "secret",
			Prompt:    &survey.Input{Message: "Client secret"},
			Transform: trans,
		},
	}

	ct := api.ClientToken{}
	if err := survey.Ask(qs, &ct); err != nil {
		return err
	}

	c.api.Token.Client = &ct
	if err := c.cfg.Save(&c.api.Token); err != nil {
		return err
	}

	view.PrintDone(cmd.OutOrStdout(), "Saved!")
	return nil
}
