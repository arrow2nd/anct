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

// ReceiveStatus : 視聴ステータスを受け取る
func ReceiveStatus(p *pflag.FlagSet, allowNoState bool) (gen.StatusState, error) {
	statusStr, _ := p.GetString("status")
	if statusStr == "" {
		// フラグで指定されていない場合、対話形式で聞く
		s, err := view.SelectStatus(allowNoState)
		if err != nil {
			return "", err
		}
		statusStr = s
	}

	return StringToStatusState(statusStr, allowNoState)
}
