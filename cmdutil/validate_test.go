package cmdutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateSeasonFormat(t *testing.T) {
	t.Run("正しい書式を判定できるか", func(t *testing.T) {
		tests := []string{
			"2005-winter",
			"2006-spring",
			"2007-summer",
			"2008-winter",
		}

		for _, tt := range tests {
			err := validateSeasonFormat(tt)
			assert.NoError(t, err, "str = %s", tt)
		}
	})

	t.Run("不正な書式を判定できるか", func(t *testing.T) {
		tests := []string{
			"0-winter",
			"20-spring",
			"200-summer",
			"20008-winter",
			"2005-WINTER",
			"2005winter",
			"2006-all",
			"",
		}

		for _, tt := range tests {
			err := validateSeasonFormat(tt)
			assert.Error(t, err, "str = %s", tt)
		}
	})
}
