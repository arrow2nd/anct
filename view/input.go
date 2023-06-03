package view

import (
	"errors"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// Confirm : 確認ダイアログ
func Confirm(m string) (bool, error) {
	prompt := &survey.Confirm{
		Message: m,
		Default: true,
	}

	ok := false
	if err := survey.AskOne(prompt, &ok); err != nil {
		return false, err
	}

	return ok, nil
}

// InputText : テキストを入力
func InputText(m string, allowEmpty bool) (string, error) {
	prompt := &survey.Input{
		Message: m,
	}

	var opt survey.AskOpt = nil

	if !allowEmpty {
		opt = survey.WithValidator(func(ans interface{}) error {
			if str, ok := ans.(string); !ok || len(str) == 0 {
				return errors.New("please enter text")
			}
			return nil
		})
	}

	text := ""
	err := survey.AskOne(prompt, &text, opt)

	return text, err
}

// InputTextInEditor : エディタを開いて文字を入力
func InputTextInEditor(m string) (string, error) {
	prompt := &survey.Editor{
		Message:  m,
		FileName: "*.txt",
	}

	text := ""
	if err := survey.AskOne(prompt, &text); err != nil {
		return "", err
	}

	return strings.TrimSpace(text), nil
}
