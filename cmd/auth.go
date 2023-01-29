package cmd

import (
	"errors"
	"fmt"

	"github.com/arrow2nd/anct/api"
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

func (c *Cmd) newAuthCmd() *cobra.Command {
	auth := &cobra.Command{
		Use:   "auth",
		Short: "Authentication anct with Annict",
		Args:  cobra.NoArgs,
	}

	login := &cobra.Command{
		Use:   "login",
		Short: "Authentication with Annict",
		Args:  cobra.NoArgs,
		RunE:  execLogin,
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

func execLogin(cmd *cobra.Command, args []string) error {
	url, err := api.GetAuthorizeURL()
	if err != nil {
		return err
	}

	fmt.Printf(`%s
Please access the following URL and enter the code displayed after authentication.
URL: %s
`, logo, url)

	code, err := inputCode()
	if err != nil {
		return err
	}

	cred, err := api.GetToken(code)
	if err != nil {
		return err
	}

	// TODO: credをファイルに保存
	fmt.Println(cred)

	return nil
}
