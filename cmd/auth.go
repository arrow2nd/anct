package cmd

import "github.com/spf13/cobra"

func (c *Cmd) newAuthCmd() *cobra.Command {
	auth := &cobra.Command{
		Use:   "auth",
		Short: "Authentication annict-for-term with Annict",
		Args:  cobra.NoArgs,
	}

	login := &cobra.Command{
		Use:   "login",
		Short: "Authentication with Annict",
		Args:  cobra.NoArgs,
	}

	logout := &cobra.Command{
		Use:   "logout",
		Short: "Log out of annict-for-term",
		Args:  cobra.NoArgs,
	}

	auth.AddCommand(
		login,
		logout,
	)

	return auth
}
