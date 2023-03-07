package cmdutil

import (
	"fmt"
	"regexp"
)

// validateSeasonFormat : シーズン指定の書式が正しいか
func validateSeasonFormat(s string) error {
	r := regexp.MustCompile(`^\d{4}-(spring|summer|autumn|winter)$`)

	if !r.MatchString(s) {
		return fmt.Errorf("incorrect season format (%s)", s)
	}

	return nil
}
