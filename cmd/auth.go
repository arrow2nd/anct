package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/arrow2nd/anct/config"
	"github.com/arrow2nd/anct/view"
	"github.com/spf13/cobra"
)

func (c *Command) newCmdAuth() *cobra.Command {
	auth := &cobra.Command{
		Use:   "auth",
		Short: "Authentication anct with Annict",
		Args:  cobra.NoArgs,
	}

	login := &cobra.Command{
		Use:   "login",
		Short: "Authentication with Annict",
		Args:  cobra.NoArgs,
		RunE:  c.loginRun,
	}

	logout := &cobra.Command{
		Use:   "logout",
		Short: "Log out of anct",
		Args:  cobra.NoArgs,
		RunE:  c.logoutRun,
	}

	auth.AddCommand(
		login,
		logout,
	)

	return auth
}

func (c *Command) loginRun(cmd *cobra.Command, args []string) error {
	url, err := c.api.CreateAuthorizeURL()
	if err != nil {
		return err
	}

	view.PrintLogo(os.Stdout)
	view.PrintAuthURL(os.Stdout, url)

	code, err := view.InputAuthCode()
	if err != nil {
		return err
	}

	if err := c.api.UpdateUserToken(code); err != nil {
		return err
	}

	return config.Save(&c.api.Token)
}

func (c *Command) logoutRun(cmd *cobra.Command, arg []string) error {
	isLogout := false
	prompt := &survey.Confirm{
		Message: "Do you want to log out?",
	}

	if err := survey.AskOne(prompt, &isLogout); err != nil {
		return err
	}

	if !isLogout {
		view.PrintCanceled(os.Stderr)
		return nil
	}

	if err := c.api.Token.Revoke(); err != nil {
		return fmt.Errorf("failed to revoke access token: %w", err)
	}

	return config.Save(&c.api.Token)
}
