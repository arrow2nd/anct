package view

import (
	"bytes"
	"encoding/base64"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapPassthrough(t *testing.T) {
	// APC全体を\x1bPtmux;...\x1b\\で包み、内部のESCを二重化する
	in := []byte("\x1b_Gx;y\x1b\\")
	want := "\x1bPtmux;\x1b\x1b_Gx;y\x1b\x1b\\\x1b\\"

	assert.Equal(t, want, string(wrapPassthrough(in)))
}

func TestWritePlaceholders(t *testing.T) {
	ph := string(placeholderRune)
	d0 := string(rowColumnDiacritics[0]) // 行/列 0
	d1 := string(rowColumnDiacritics[1]) // 行/列 1

	buf := &bytes.Buffer{}
	if err := writePlaceholders(buf, 1, 2, 2); err != nil {
		t.Fatal(err)
	}

	// id=1 → 前景色 (0,0,1)。各セルが placeholder + 行diacritic + 列diacritic
	want := "\x1b[38;2;0;0;1m" + ph + d0 + d0 + ph + d0 + d1 + "\x1b[39m\n" +
		"\x1b[38;2;0;0;1m" + ph + d1 + d0 + ph + d1 + d1 + "\x1b[39m\n"

	assert.Equal(t, want, buf.String())
}

func TestWriteKGPTransmit_SingleChunk(t *testing.T) {
	payload := []byte("hello")
	b64 := base64.StdEncoding.EncodeToString(payload)
	raw := "\x1b_Ga=T,U=1,i=1,f=100,t=d,c=3,r=2,q=2,m=0;" + b64 + "\x1b\\"

	t.Run("tmux内: APCをpassthroughで包む", func(t *testing.T) {
		buf := &bytes.Buffer{}
		if err := writeKGPTransmit(buf, payload, 1, 3, 2, true); err != nil {
			t.Fatal(err)
		}
		// passthroughでESCが二重化され前後が\x1bPtmux;...\x1b\\で包まれる
		want := "\x1bPtmux;\x1b\x1b_Ga=T,U=1,i=1,f=100,t=d,c=3,r=2,q=2,m=0;" + b64 + "\x1b\x1b\\\x1b\\"
		assert.Equal(t, want, buf.String())
	})

	t.Run("tmux外: 生のAPCをそのまま出力", func(t *testing.T) {
		buf := &bytes.Buffer{}
		if err := writeKGPTransmit(buf, payload, 1, 3, 2, false); err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, raw, buf.String())
	})
}

func TestWriteKGPTransmit_MultiChunk(t *testing.T) {
	// 4096(chunkSize)を超えるデータは複数チャンクに分割される
	payload := bytes.Repeat([]byte("a"), 4000) // base64で約5336文字 → 2チャンク

	t.Run("tmux内: 各チャンクをpassthroughで包む", func(t *testing.T) {
		buf := &bytes.Buffer{}
		if err := writeKGPTransmit(buf, payload, 1, 3, 2, true); err != nil {
			t.Fatal(err)
		}
		out := buf.String()
		assert.Equal(t, 2, strings.Count(out, "\x1bPtmux;"))
		assert.Contains(t, out, "a=T,U=1,i=1,f=100,t=d,c=3,r=2,q=2,m=1;")
		assert.Contains(t, out, "\x1b\x1b_Gm=0;")
	})

	t.Run("tmux外: passthroughなしで2チャンク", func(t *testing.T) {
		buf := &bytes.Buffer{}
		if err := writeKGPTransmit(buf, payload, 1, 3, 2, false); err != nil {
			t.Fatal(err)
		}
		out := buf.String()
		assert.Equal(t, 0, strings.Count(out, "\x1bPtmux;"))
		assert.Equal(t, 2, strings.Count(out, "\x1b_G"))
		// 1チャンク目はcontrol data付きで継続(m=1)、2チャンク目はm=0のみ
		assert.Contains(t, out, "a=T,U=1,i=1,f=100,t=d,c=3,r=2,q=2,m=1;")
		assert.Contains(t, out, "\x1b_Gm=0;")
	})
}

func TestRowColumnDiacritics(t *testing.T) {
	// kittyのrowcolumn-diacritics.txtは297エントリ。
	// calcCellsがこの長さを上限に使うため、欠落していないことを保証する
	assert.Len(t, rowColumnDiacritics, 297)
	assert.Equal(t, rune(0x0305), rowColumnDiacritics[0])
	assert.Equal(t, rune(0x1D244), rowColumnDiacritics[296])
}
