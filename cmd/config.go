package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/arrow2nd/anct/api"
	"github.com/arrow2nd/anct/config"
	"github.com/arrow2nd/anct/view"
	"github.com/spf13/cobra"
)

func (c *Command) newCmdConfig() *cobra.Command {
	config := &cobra.Command{
		Use:     "config",
		Short:   "Edit the configuration file",
		Example: "  anct config",
	}

	cred := &cobra.Command{
		Use:     "apikey",
		Short:   "Change the API key used to connect to Annict",
		Example: "  anct config apikey",
		RunE:    c.conifgRun,
	}

	config.AddCommand(cred)

	return config
}

func (c *Command) conifgRun(cmd *cobra.Command, args []string) error {
	qs := []*survey.Question{
		{
			Name:   "id",
			Prompt: &survey.Input{Message: "Client ID"},
		},
		{
			Name:   "secret",
			Prompt: &survey.Input{Message: "Client secret"},
		},
	}

	token := api.ClientToken{}
	if err := survey.Ask(qs, &token); err != nil {
		return err
	}

	c.api.Token.Client = &token
	if err := config.Save(&c.api.Token); err != nil {
		return err
	}

	view.PrintDone(cmd.OutOrStdout(), "Saved!")
	return nil
}
