package cmd

import (
	"fmt"

	"github.com/arrow2nd/anct/cmdutil"
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
		Short: "Authentication with a Annict",
		Args:  cobra.NoArgs,
		RunE:  c.loginRun,
	}

	logout := &cobra.Command{
		Use:   "logout",
		Short: "Log out of a Annict",
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

	view.PrintLogo(cmd.OutOrStdout())
	view.PrintAuthURL(cmd.OutOrStdout(), url)

	code, err := view.InputText("Code", false)
	if err != nil {
		return err
	}

	code = cmdutil.StripWhiteSpace(code)
	if err := c.api.UpdateUserToken(code); err != nil {
		return err
	}

	if err := c.cfg.Save(&c.api.Token); err != nil {
		return err
	}

	view.PrintDone(cmd.OutOrStdout(), "Logged in!")
	return nil
}

func (c *Command) logoutRun(cmd *cobra.Command, arg []string) error {
	logout, err := view.Confirm("Are you sure you want to log out?")
	if err != nil {
		return err
	}

	if !logout {
		view.PrintCanceled(cmd.ErrOrStderr())
		return nil
	}

	if err := c.api.Token.Revoke(); err != nil {
		return fmt.Errorf("failed to revoke access token: %w", err)
	}

	if err := c.cfg.Save(&c.api.Token); err != nil {
		return err
	}

	view.PrintDone(cmd.ErrOrStderr(), "Logged out")
	return nil
}
