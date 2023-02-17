package view

import (
	"fmt"
	"image"
	"io"
	"net/http"

	"github.com/disintegration/imaging"
	"github.com/mattn/go-sixel"
)

// printImage : 画像を出力 (sixelを使用)
func printImage(w io.Writer, URL string) error {
	img, err := fetchImage(URL, 400)
	if err != nil {
		return fmt.Errorf("failed fetch image: %w", err)
	}

	// 画像の前に空行を挿入
	fmt.Fprintln(w)

	if err := sixel.NewEncoder(w).Encode(img); err != nil {
		return fmt.Errorf("failed print image: %w", err)
	}

	return nil
}

// fetchImage : 画像を取得
func fetchImage(URL string, width int) (image.Image, error) {
	res, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)
	img = imaging.Resize(img, width, 0, imaging.Lanczos)

	return img, err
}
