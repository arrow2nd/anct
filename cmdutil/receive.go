package cmdutil

import (
	"io"
	"os"
	"strings"
	"syscall"

	"github.com/arrow2nd/anct/gen"
	"github.com/arrow2nd/anct/view"
	"github.com/spf13/pflag"
	"golang.org/x/term"
)

// receiveQuery : クエリを受け取る
func receiveQuery(m string, args []string, useEditor, allowEmpty bool) (string, error) {
	// 引数
	query := strings.Join(args, " ")
	if query != "" {
		return query, nil
	}

	// 標準入力
	if !term.IsTerminal(int(syscall.Stdin)) {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", err
		}

		return string(stdin), nil
	}

	// プロンプト or エディタ
	return ReceiveText(m, useEditor, allowEmpty)
}

// ReceiveText : プロンプト・外部エディタから文字列を受け取る
func ReceiveText(m string, useEditor, allowEmpty bool) (string, error) {
	if useEditor {
		return view.InputTextInEditor(m)
	}

	return view.InputText(m, allowEmpty)
}

// ReceiveRating : 評価を受け取る
func ReceiveRating(p *pflag.FlagSet, flagName string) (gen.RatingState, error) {
	rating, _ := p.GetString(flagName)

	// 指定されていない場合対話形式で聞く
	if rating == "" {
		m := StringToUpperFirstLetter(flagName)
		r, err := view.SelectRating(m)
		if err != nil {
			return "", err
		}

		rating = r
	}

	return StringToRatingState(rating)
}

// ReceiveBody : フラグから Body を受け取る
func ReceiveBody(p *pflag.FlagSet, flagName string) (string, error) {
	text, _ := p.GetString(flagName)

	// 指定されていなければエディタを開く
	if text == "" {
		return view.InputTextInEditor(StringToUpperFirstLetter(flagName))
	}

	return text, nil
}
