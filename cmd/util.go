package cmd

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/arrow2nd/anct/gen"
	"github.com/spf13/pflag"
)

// setLimitFlag : limit フラグを設定
func setLimitFlag(p *pflag.FlagSet) {
	p.Int64P("limit", "l", 30, "Maximum number of results to fetch")
}

// setEditerFlag : editor フラグを設定
func setEditerFlag(p *pflag.FlagSet) {
	p.BoolP("editor", "e", false, "Use an external editor to enter keyword")
}

// checkSeasonFormat : シーズン指定の書式が正しいか
func checkSeasonFormat(s string) error {
	r := regexp.MustCompile(`\d{4}-(spring|summer|autumn|winter)`)

	if !r.MatchString(s) {
		return fmt.Errorf("incorrect season format (%s)", s)
	}

	return nil
}

// toWorkID : 文字列を Work IDに 変換
func toWorkID(s string) (int64, error) {
	workID, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("work id must be numeric")
	}

	if workID <= 0 {
		return 0, errors.New("work id must be greater than or equal to 1")
	}

	return int64(workID), nil
}

// toStatusState : 文字列を視聴ステータスに変換
func toStatusState(s string, allowNoState bool) (gen.StatusState, error) {
	for _, status := range gen.AllStatusState {
		if !allowNoState && status == gen.StatusStateNoState {
			continue
		}
		if string(status) == strings.ToUpper(s) {
			return status, nil
		}
	}

	return "", fmt.Errorf("incorrect status (%s)", s)
}
