package view

import (
	"io"
	"time"

	"github.com/briandowns/spinner"
)

// NewSpinner : スピナーを作成
func NewSpinner(w io.Writer, m string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)

	s.Writer = w
	s.Suffix = "  " + m
	s.Color("cyan")

	return s
}
