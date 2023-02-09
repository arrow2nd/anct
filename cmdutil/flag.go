package cmdutil

import "github.com/spf13/pflag"

// SetLimitFlag : limit フラグを設定
func SetLimitFlag(p *pflag.FlagSet) {
	p.Int64P("limit", "l", 30, "Maximum number of results to fetch")
}

// SetEditerFlag : editor フラグを設定
func SetEditerFlag(p *pflag.FlagSet) {
	p.BoolP("editor", "e", false, "Use an external editor to enter keyword")
}
