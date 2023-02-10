package cmdutil

import (
	"errors"
	"io"
	"os"
	"regexp"
	"strings"
	"syscall"

	"github.com/arrow2nd/anct/gen"
	"github.com/arrow2nd/anct/view"
	"github.com/spf13/pflag"
	"golang.org/x/term"
)

// ReceiveQuery : キーワードの入力を受け取る
func ReceiveQuery(args []string, useEditor bool, allowEmpty bool) (string, error) {
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

// ReceiveCommonFlags : 全体共通フラグの値を受け取る
func ReceiveCommonFlags(p *pflag.FlagSet) (bool, int64) {
	useEditor, _ := p.GetBool("editor")
	limit, _ := p.GetInt64("limit")
	return useEditor, limit
}

// ReceiveSearchCommonFlags : 検索共通フラグの値を受け取る
func ReceiveSearchCommonFlags(p *pflag.FlagSet) ([]gen.StatusState, []string, error) {
	// シーズン指定の書式をチェック
	seasons, _ := p.GetStringSlice("seasons")
	for _, s := range seasons {
		if err := ValidateSeasonFormat(s); err != nil {
			return nil, nil, err
		}
	}

	// ライブラリの視聴ステータスを取得
	stateStrs, _ := p.GetStringSlice("library")
	if len(stateStrs) == 0 {
		return nil, seasons, nil
	}

	// 視聴ステータス文字列を変換
	states := []gen.StatusState{}
	for _, stateStr := range stateStrs {
		s, err := StringToStatusState(stateStr, false)
		if err != nil {
			return nil, nil, err
		}
		states = append(states, s)
	}

	return states, seasons, nil
}
