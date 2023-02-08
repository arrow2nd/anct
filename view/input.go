package view

import (
	"errors"
	"io"
	"os"
	"regexp"
	"strings"
	"syscall"

	"github.com/AlecAivazis/survey/v2"
	"golang.org/x/term"
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

// Receivekeyword : キーワードの入力を受け取る
func Receivekeyword(args []string, useEditor bool, allowEmpty bool) (string, error) {
	keyword := strings.Join(args, " ")

	// 標準入力を受け取る
	if keyword == "" && !term.IsTerminal(int(syscall.Stdin)) {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", err
		}
		keyword = strings.TrimSpace(string(stdin))
	}

	// エディタを起動
	if keyword == "" && useEditor {
		s, err := InputTextInEditor("Enter search keyword")
		if err != nil {
			return "", err
		}
		keyword = s
	}

	if keyword == "" && !allowEmpty {
		return "", errors.New("please enter keywords")
	}

	// 全ての空白文字を半角スペースに置換
	r := regexp.MustCompile(`\s`)
	return r.ReplaceAllString(keyword, " "), nil
}
