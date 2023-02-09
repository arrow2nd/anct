package view

import (
	"errors"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// Confirm : 確認ダイアログ
func Confirm(m string) (bool, error) {
	result := false
	prompt := &survey.Confirm{
		Message: m,
	}

	if err := survey.AskOne(prompt, &result); err != nil {
		return false, err
	}

	return result, nil
}

// InputText : テキストを入力
func InputText(m string) (string, error) {
	prompt := &survey.Input{
		Message: m,
	}

	validator := func(ans interface{}) error {
		if str, ok := ans.(string); !ok || len(str) == 0 {
			return errors.New("please enter text")
		}
		return nil
	}

	text := ""
	err := survey.AskOne(prompt, &text, survey.WithValidator(validator))

	return text, err
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
