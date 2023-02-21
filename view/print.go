package view

import (
	"fmt"
	"io"
)

// PrintDone : 完了表示を出力
func PrintDone(w io.Writer, s string) {
	fmt.Fprintf(w, "👌 %s\n", s)
}

// PrintCanceled : キャンセル表示を出力
func PrintCanceled(w io.Writer) {
	fmt.Fprintln(w, "❌ Canceled")
}

// PrintLogo : ロゴを出力
func PrintLogo(w io.Writer) {
	logo := `
   ________  ________  ________  ________ 
  /        \/    /   \/        \/        \
 /         /         /         /        _/
/         /         /       --//       /  
\___/____/\__/_____/\________/ \______/

`

	fmt.Fprint(w, logo)
}

// PrintAuthURL : 認証URLを出力
func PrintAuthURL(w io.Writer, u string) {
	temp := "📺 Please access the following URL and enter the code displayed after authentication.\n\n%s\n\n"
	fmt.Fprintf(w, temp, u)
}
