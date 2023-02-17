package view

import (
	"io"
	"time"

	"github.com/briandowns/spinner"
)

// SpinnerStart : スピナーを作成 & 開始
func SpinnerStart(w io.Writer, m string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)

	s.Writer = w
	s.Suffix = "  " + m
	s.Color("cyan")

	s.Start()
	return s
}
