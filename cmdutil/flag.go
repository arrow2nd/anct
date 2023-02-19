package cmdutil

import (
	"github.com/arrow2nd/anct/gen"
	"github.com/spf13/pflag"
)

// SetCommonFlags : 全体共通フラグを設定
func SetCommonFlags(p *pflag.FlagSet) {
	p.BoolP("editor", "e", false, "use an external editor to enter text")
	p.Int64P("limit", "l", 30, "maximum number of results to fetch")
}

// SetSearchFlags : 検索関連フラグを設定
func SetSearchFlags(p *pflag.FlagSet) {
	SetCommonFlags(p)

	p.StringSliceP(
		"seasons",
		"S",
		[]string{},
		"retrieve works for a given season: YYYY-{spring|summer|autumn|winter}",
	)

	p.StringSliceP("library",
		"L",
		[]string{},
		"search within the library: {wanna_watch|watching|watched|on_hold|stop_watching}",
	)
}

// GetCommonFlags : 共通フラグの内容を取得
func GetCommonFlags(p *pflag.FlagSet) (bool, int64) {
	useEditor, _ := p.GetBool("editor")
	limit, _ := p.GetInt64("limit")

	return useEditor, limit
}

// getAllSearchFlags : 全て検索フラグの内容を取得
func getAllSearchFlags(p *pflag.FlagSet) ([]gen.StatusState, []string, int64, bool, error) {
	seasons, _ := p.GetStringSlice("seasons")

	// シーズン指定の書式をチェック
	for _, s := range seasons {
		if err := validateSeasonFormat(s); err != nil {
			return nil, nil, 0, false, err
		}
	}

	stateStrs, _ := p.GetStringSlice("library")
	states := []gen.StatusState{}

	// ライブラリの視聴ステータス文字列を変換
	for _, stateStr := range stateStrs {
		s, err := ConvertToStatusState(stateStr, false)
		if err != nil {
			return nil, nil, 0, false, err
		}

		states = append(states, s)
	}

	useEditor, limit := GetCommonFlags(p)

	return states, seasons, limit, useEditor, nil
}
