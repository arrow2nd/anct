package view

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"net/http"
	"os"

	"github.com/disintegration/imaging"
	"github.com/mattn/go-sixel"
)

// inTmux : tmuxセッション内で動作しているかをチェック
func inTmux() bool {
	return os.Getenv("TMUX") != ""
}

// checkKitty : kitty画像プロトコルがサポートされているかをチェック
func checkKitty() bool {
	if os.Getenv("KITTY_WINDOW_ID") != "" {
		return true
	}

	if os.Getenv("TERM_PROGRAM") == "ghostty" {
		return true
	}

	// tmux内ではTERM_PROGRAMがtmuxに上書きされるため、Ghostty固有の環境変数で判定
	if os.Getenv("GHOSTTY_RESOURCES_DIR") != "" {
		return true
	}

	return false
}

// printImage : 画像を出力 (kitty画像プロトコルまたはsixel)
func printImage(w io.Writer, img image.Image) error {
	// Kitty Graphics Protocol対応端末ではUnicode placeholder方式で出力する。
	// 直接配置はtmux内で画面に焼き付くため、tmux内外を問わずplaceholderに統一する。
	if checkKitty() {
		return printKittyPlaceholder(w, img)
	}

	if err := sixel.NewEncoder(w).Encode(img); err != nil {
		return fmt.Errorf("failed to print image with sixel: %w", err)
	}

	return nil
}

// wrapPassthrough : 1つのエスケープシーケンスをtmuxのDCS passthroughで包む。
// ペイロード内のESC(0x1b)をすべて二重化しないとtmux内で壊れる。
func wrapPassthrough(seq []byte) []byte {
	doubled := bytes.ReplaceAll(seq, []byte("\x1b"), []byte("\x1b\x1b"))

	var b bytes.Buffer
	b.WriteString("\x1bPtmux;")
	b.Write(doubled)
	b.WriteString("\x1b\\")

	return b.Bytes()
}

// fetchImage : 画像を取得
func fetchImage(URL string, width int) (image.Image, error) {
	res, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	if code := res.StatusCode; code != http.StatusOK {
		return nil, fmt.Errorf("could not retrieve (status: %d)", code)
	}

	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)
	if err != nil {
		return nil, err
	}

	return imaging.Resize(img, width, 0, imaging.Lanczos), nil
}
