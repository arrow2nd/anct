package view

import (
	"fmt"
	"io"
)

// PrintCanceled : キャンセル表示を出力
func PrintCanceled(w io.Writer) {
	fmt.Fprintln(w, "Canceled")
}

// PrintLogo : ロゴを出力
func PrintLogo(w io.Writer) {
	fmt.Fprint(w, `
   ________  ________  ________  ________ 
  /        \/    /   \/        \/        \
 /         /         /         /        _/
/         /         /       --//       /  
\___/____/\__/_____/\________/ \______/
         -- Unofficial CLI Client of Annict
`)
}

// PrintAuthURL : 認証URLを出力
func PrintAuthURL(w io.Writer, u string) {
	fmt.Fprintf(w, `Please access the following URL and enter the code displayed after authentication.
> %s

`, u)
}
