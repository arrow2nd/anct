package view

import (
	"fmt"
	"image"
	"io"
	"net/http"
	"os"

	"github.com/disintegration/imaging"
	"github.com/dolmen-go/kittyimg"
	"github.com/mattn/go-sixel"
)

// checkKitty : kitty画像プロトコルがサポートされているかをチェック
func checkKitty() bool {
	if os.Getenv("KITTY_WINDOW_ID") != "" {
		return true
	}

	if os.Getenv("TERM_PROGRAM") == "ghostty" {
		return true
	}

	return false
}

// printImage : 画像を出力 (kitty画像プロトコルまたはsixelを使用)
func printImage(w io.Writer, img image.Image) error {
	isKitty := checkKitty()

	if isKitty {
		// Kittyターミナルの場合はkitty画像プロトコルを使用
		if err := kittyimg.Fprint(w, img); err != nil {
			return fmt.Errorf("failed to print image with kitty protocol: %w", err)
		}
	} else {
		// それ以外はsixelを使用
		if err := sixel.NewEncoder(w).Encode(img); err != nil {
			return fmt.Errorf("failed to print image with sixel: %w", err)
		}
	}

	return nil
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
