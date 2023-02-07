package view

import (
	"errors"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// InputAuthCode : 認証コードの入力
func InputAuthCode() (string, error) {
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

// InputTextInEditor : エディタを開いて文字を入力
func InputTextInEditor(m string) (string, error) {
	s := ""

	prompt := &survey.Editor{
		Message:  m,
		FileName: "*.txt",
	}

	if err := survey.AskOne(prompt, &s); err != nil {
		return "", err
	}

	return strings.TrimSpace(s), nil
}
