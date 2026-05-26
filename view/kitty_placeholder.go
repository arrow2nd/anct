package view

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"io"
	"os"

	"github.com/disintegration/imaging"
	"golang.org/x/sys/unix"
)

// imageID : 表示する画像のID。infoは単一画像のみ表示するため固定値でよい
// (毎回 a=T で同じIDに上書き転送されるため衝突しない)。
const imageID = 1

// printKittyPlaceholder : Unicode placeholder方式でKGP画像を出力する。
// 直接配置(a=T)はtmux内で画面に焼き付いて消せなくなるため、画像をテキストセルとして
// 配置できるplaceholder方式に統一する。tmux内ではエスケープをpassthroughで包む。
func printKittyPlaceholder(w io.Writer, img image.Image) error {
	cols, rows := calcCells(img)

	// f=100(PNG)で転送する。RGBA生データより小さく、kitty側でデコードされる
	var png bytes.Buffer
	if err := imaging.Encode(&png, img, imaging.PNG); err != nil {
		return fmt.Errorf("failed to encode png: %w", err)
	}

	if err := writeKGPTransmit(w, png.Bytes(), imageID, cols, rows, inTmux()); err != nil {
		return fmt.Errorf("failed to transmit kitty image: %w", err)
	}

	if err := writePlaceholders(w, imageID, cols, rows); err != nil {
		return fmt.Errorf("failed to write placeholders: %w", err)
	}

	return nil
}

// calcCells : 画像を表示するセル数(列・行)を算出する。
// セルのピクセルサイズから元画像と同等の見た目になるよう求めるが、
// tmux内ではピクセルサイズが取得できない(0が返る)ことがあるためフォールバックする。
func calcCells(img image.Image) (cols, rows int) {
	const fallbackCellW, fallbackCellH = 10, 20

	bounds := img.Bounds()
	imgW, imgH := bounds.Dx(), bounds.Dy()

	cellW, cellH := fallbackCellW, fallbackCellH
	termCols := 0

	if ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ); err == nil && ws.Col > 0 && ws.Row > 0 {
		termCols = int(ws.Col)
		if ws.Xpixel > 0 && ws.Ypixel > 0 {
			cellW = int(ws.Xpixel) / int(ws.Col)
			cellH = int(ws.Ypixel) / int(ws.Row)
		}
	}

	if cellW <= 0 {
		cellW = fallbackCellW
	}
	if cellH <= 0 {
		cellH = fallbackCellH
	}

	cols = max(imgW/cellW, 1)
	rows = max(imgH/cellH, 1)

	// 端末幅を超える場合はアスペクト比を維持して縮小
	if termCols > 0 && cols > termCols {
		rows = max(rows*termCols/cols, 1)
		cols = termCols
	}

	// 行・列番号はdiacriticテーブルの範囲内に収める必要がある
	maxN := len(rowColumnDiacritics)
	if cols > maxN {
		rows = max(rows*maxN/cols, 1)
		cols = maxN
	}
	if rows > maxN {
		rows = maxN
	}

	return cols, rows
}

// writeKGPTransmit : PNGをbase64化し、4096文字ごとにチャンク分割して転送する。
// a=T,U=1で転送と仮想配置を同時に行う。wrapTmuxがtrueのとき、各チャンクのAPCを
// tmux passthroughで包む(tmux内ではそうしないとエスケープが捨てられる)。
func writeKGPTransmit(w io.Writer, png []byte, id, cols, rows int, wrapTmux bool) error {
	const chunkSize = 4096

	b64 := base64.StdEncoding.EncodeToString(png)
	first := true

	for len(b64) > 0 {
		n := min(chunkSize, len(b64))
		chunk := b64[:n]
		b64 = b64[n:]

		// 後続チャンクの有無を m=1/m=0 で示す
		more := 0
		if len(b64) > 0 {
			more = 1
		}

		var seq bytes.Buffer
		seq.WriteString("\x1b_G")
		if first {
			fmt.Fprintf(&seq, "a=T,U=1,i=%d,f=100,t=d,c=%d,r=%d,q=2,m=%d;", id, cols, rows, more)
			first = false
		} else {
			// 2チャンク目以降はcontrol dataを省略しmのみ指定する
			fmt.Fprintf(&seq, "m=%d;", more)
		}
		seq.WriteString(chunk)
		seq.WriteString("\x1b\\")

		payload := seq.Bytes()
		if wrapTmux {
			payload = wrapPassthrough(payload)
		}

		if _, err := w.Write(payload); err != nil {
			return err
		}
	}

	return nil
}

// writePlaceholders : placeholder文字(U+10EEEE)をcols×rowsのグリッドで出力する。
// 前景色でimage IDをエンコードし、各セルに行・列のdiacriticを付ける。
// placeholder文字は通常テキストなのでtmux passthroughは不要。
func writePlaceholders(w io.Writer, id, cols, rows int) error {
	// 24bit前景色でimage IDをエンコード
	r := (id >> 16) & 0xff
	g := (id >> 8) & 0xff
	b := id & 0xff

	var out bytes.Buffer
	for row := 0; row < rows; row++ {
		fmt.Fprintf(&out, "\x1b[38;2;%d;%d;%dm", r, g, b)
		for col := 0; col < cols; col++ {
			out.WriteRune(placeholderRune)
			out.WriteRune(rowColumnDiacritics[row])
			out.WriteRune(rowColumnDiacritics[col])
		}
		out.WriteString("\x1b[39m\n")
	}

	_, err := w.Write(out.Bytes())
	return err
}
