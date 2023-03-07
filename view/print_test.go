package view_test

import (
	"bytes"
	"testing"

	"github.com/arrow2nd/anct/view"
	"github.com/stretchr/testify/assert"
)

func TestPrintDone(t *testing.T) {
	buf := &bytes.Buffer{}
	want := "test"

	view.PrintDone(buf, want)
	assert.Contains(t, buf.String(), want)
}

func TestPrintCanceled(t *testing.T) {
	buf := &bytes.Buffer{}
	view.PrintCanceled(buf)

	want := "Canceled"
	assert.Contains(t, buf.String(), want)
}

func TestPrintAuthURL(t *testing.T) {
	buf := &bytes.Buffer{}
	want := "https://example.com"

	view.PrintAuthURL(buf, want)
	assert.Contains(t, buf.String(), want)
}
