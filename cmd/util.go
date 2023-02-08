package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"syscall"

	"github.com/arrow2nd/anct/view"
	"github.com/spf13/pflag"
	"golang.org/x/term"
)

// setLimitFlag : limit フラグを設定
func setLimitFlag(p *pflag.FlagSet) {
	p.Int64P("limit", "l", 30, "Maximum number of results to fetch")
}

// setEditerFlag : editor フラグを設定
func setEditerFlag(p *pflag.FlagSet) {
	p.BoolP("editor", "e", false, "Use an external editor to enter keyword")
}

// receivekeyword : キーワードの入力を受け取る
func receivekeyword(args []string, useEditor bool, allowEmpty bool) (string, error) {
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
		s, err := view.InputTextInEditor("Enter search keyword")
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

// checkSeasonFormat : シーズン指定の書式が正しいか
func checkSeasonFormat(s string) error {
	r := regexp.MustCompile(`\d{4}-(spring|summer|autumn|winter)`)

	if !r.MatchString(s) {
		return fmt.Errorf("incorrect season format (%s)", s)
	}

	return nil
}
