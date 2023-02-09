package cmdutil

import (
	"fmt"
	"regexp"
)

// ValidateSeasonFormat : シーズン指定の書式が正しいか
func ValidateSeasonFormat(s string) error {
	r := regexp.MustCompile(`\d{4}-(spring|summer|autumn|winter)`)

	if !r.MatchString(s) {
		return fmt.Errorf("incorrect season format (%s)", s)
	}

	return nil
}
