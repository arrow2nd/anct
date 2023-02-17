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

// ReceiveAllSearchFlags : 全ての検索フラグの値を受け取る
func ReceiveAllSearchFlags(p *pflag.FlagSet) ([]gen.StatusState, []string, int64, bool, error) {
	// シーズン指定の書式をチェック
	seasons, _ := p.GetStringSlice("seasons")
	for _, s := range seasons {
		if err := validateSeasonFormat(s); err != nil {
			return nil, nil, 0, false, err
		}
	}

	// ライブラリの視聴ステータス文字列を変換
	stateStrs, _ := p.GetStringSlice("library")
	states := []gen.StatusState{}
	for _, stateStr := range stateStrs {
		s, err := StringToStatusState(stateStr, false)
		if err != nil {
			return nil, nil, 0, false, err
		}
		states = append(states, s)
	}

	useEditor, _ := p.GetBool("editor")
	limit, _ := p.GetInt64("limit")
	return states, seasons, limit, useEditor, nil
}

// ReceiveQuery : クエリの入力を受け取る
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

// ReceiveRating : 評価を受け取る
func ReceiveRating(p *pflag.FlagSet) (gen.RatingState, error) {
	rating, _ := p.GetString("rating")

	if rating == "" {
		// 指定されていない場合対話形式で聞く
		r, err := view.SelectRating()
		if err != nil {
			return "", err
		}
		rating = r
	}

	return StringToRatingState(rating)
}
