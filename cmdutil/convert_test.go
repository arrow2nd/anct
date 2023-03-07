package cmdutil

import (
	"testing"

	"github.com/arrow2nd/anct/gen"
	"github.com/stretchr/testify/assert"
)

func TestConvertToStatusState(t *testing.T) {
	t.Run("変換できるか", func(t *testing.T) {
		tests := []struct {
			str  string
			want gen.StatusState
		}{
			{
				str:  "wanna_watch",
				want: gen.StatusStateWannaWatch,
			},
			{
				str:  "watching",
				want: gen.StatusStateWatching,
			},
			{
				str:  "watched",
				want: gen.StatusStateWatched,
			},
			{
				str:  "on_hold",
				want: gen.StatusStateOnHold,
			},
			{
				str:  "stop_watching",
				want: gen.StatusStateStopWatching,
			},
			{
				str:  "no_state",
				want: gen.StatusStateNoState,
			},
			{
				str:  "WANNA_watch",
				want: gen.StatusStateWannaWatch,
			},
			{
				str:  "WATCHing",
				want: gen.StatusStateWatching,
			},
			{
				str:  "watchED",
				want: gen.StatusStateWatched,
			},
			{
				str:  "on_HOLD",
				want: gen.StatusStateOnHold,
			},
			{
				str:  "stop_WATCHing",
				want: gen.StatusStateStopWatching,
			},
			{
				str:  "No_statE",
				want: gen.StatusStateNoState,
			},
		}

		for _, tt := range tests {
			got, err := ConvertToStatusState(tt.str, true)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		}
	})

	t.Run("NO_STATEが除外できるか", func(t *testing.T) {
		_, err := ConvertToStatusState("no_state", false)
		assert.EqualError(t, err, "incorrect status (no_state)")
	})
}

func TestConvertToRatingState(t *testing.T) {
	tests := []struct {
		str  string
		want gen.RatingState
	}{
		{
			str:  "great",
			want: gen.RatingStateGreat,
		},
		{
			str:  "good",
			want: gen.RatingStateGood,
		},
		{
			str:  "average",
			want: gen.RatingStateAverage,
		},
		{
			str:  "bad",
			want: gen.RatingStateBad,
		},
		{
			str:  "GReat",
			want: gen.RatingStateGreat,
		},
		{
			str:  "goOD",
			want: gen.RatingStateGood,
		},
		{
			str:  "avERAge",
			want: gen.RatingStateAverage,
		},
		{
			str:  "bAd",
			want: gen.RatingStateBad,
		},
	}

	for _, tt := range tests {
		got, err := convertToRatingState(tt.str)
		assert.NoError(t, err)
		assert.Equal(t, tt.want, got)
	}
}

func TestConvertToUpperFirstLetter(t *testing.T) {
	want := "Asahi"
	assert.Equal(t, want, "asahi")
}

func TestStripWhiteSpace(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		{
			str:  "　Serizawa　Asahi　",
			want: "Serizawa Asahi",
		},
		{
			str:  "   Mayuzumi Fuyuko  ",
			want: "Mayuzumi Fuyuko",
		},
		{
			str: `
Izumi
Mei
`,
			want: "Izumi Mei",
		},
	}

	for _, tt := range tests {
		got := StripWhiteSpace(tt.str)
		assert.Equal(t, tt.want, got)
	}
}
