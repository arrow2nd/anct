package cmd

import (
	"errors"
	"fmt"

	"github.com/arrow2nd/anct/credencial"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const logo = `
   ________  ________  ________  ________ 
  /        \/    /   \/        \/        \
 /         /         /         /        _/
/         /         /       --//       /  
\___/____/\__/_____/\________/ \______/
         -- Unofficial CLI Client of Annict
`

func (c *App) newCmdAuth() *cobra.Command {
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
	}

	auth.AddCommand(
		login,
		logout,
	)

	return auth
}

func inputCode() (string, error) {
	prompt := promptui.Prompt{
		Label: "Code",
		Validate: func(s string) error {
			if s == "" {
				return errors.New("please enter a code")
			}
			return nil
		},
	}

	return prompt.Run()
}

func (c *App) loginRun(cmd *cobra.Command, args []string) error {
	url, err := c.client.CreateAuthorizeURL()
	if err != nil {
		return err
	}

	fmt.Printf(`%s
Please access the following URL and enter the code displayed after authentication.
> %s

`, logo, url)

	code, err := inputCode()
	if err != nil {
		return err
	}

	if err := c.client.UpdateUserToken(code); err != nil {
		return err
	}

	return credencial.Save(&c.client.Token)
}

func (a *App) logoutRun(cmd *cobra.Command, arg []string) error {
	return nil
}
