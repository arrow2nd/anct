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
func printImage(w io.Writer, img image.Image) error {
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
