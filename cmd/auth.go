package cmd

import (
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/arrow2nd/anct/config"
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

	return config.Save(&c.client.Token)
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
		fmt.Println("canceled")
		return nil
	}

	if err := c.client.Token.Revoke(); err != nil {
		return fmt.Errorf("failed to revoke access token: %w", err)
	}

	return config.Save(&c.client.Token)
}

func inputCode() (string, error) {
	prompt := &survey.Input{
		Message: "Code",
	}

	validator := func(ans interface{}) error {
		if str, ok := ans.(string); !ok || len(str) == 0 {
			return errors.New("please enter a code")
		}
		return nil
	}

	code := ""
	err := survey.AskOne(prompt, &code, survey.WithValidator(validator))

	return code, err
}
