package cmd

import (
	"io"
	"os"
	"strings"
	"syscall"

	"github.com/arrow2nd/anct/view"
	"golang.org/x/term"
)

// receivekeyword : キーワードの入力を受け取る
func receivekeyword(arg []string) (string, error) {
	// 標準入力を受け取る
	if len(arg) == 0 && !term.IsTerminal(int(syscall.Stdin)) {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(stdin)), nil
	}

	// 引数がなければエディタを起動
	if len(arg) == 0 {
		return view.InputTextInEditor("Enter search keyword")
	}

	// すべての引数を連結
	return strings.Join(arg, " "), nil
}
