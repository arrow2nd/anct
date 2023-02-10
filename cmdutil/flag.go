package cmdutil

import "github.com/spf13/pflag"

// SetCommonFlags : 全体共通フラグを設定
func SetCommonFlags(p *pflag.FlagSet) {
	p.BoolP("editor", "e", false, "Use an external editor to enter keyword")
	p.Int64P("limit", "l", 30, "Maximum number of results to fetch")
}

// SetSearchFlags : 検索関連フラグを設定
func SetSearchFlags(p *pflag.FlagSet) {
	SetCommonFlags(p)
	p.StringSliceP("seasons", "s", []string{}, "Retrieve works for a given season: YYYY-{spring|summer|autumn|winter}")
	p.StringSliceP("library", "", []string{}, "Search within the library: {wanna_watch|watching|watched|on_hold|stop_watching}")
}
