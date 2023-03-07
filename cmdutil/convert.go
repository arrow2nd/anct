package cmdutil

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/arrow2nd/anct/gen"
)

// ConvertToStatusState : 文字列を StatusState に変換
func ConvertToStatusState(s string, allowNoState bool) (gen.StatusState, error) {
	for _, status := range gen.AllStatusState {
		if !allowNoState && status == gen.StatusStateNoState {
			continue
		}

		if status.String() == strings.ToUpper(s) {
			return status, nil
		}
	}

	return "", fmt.Errorf("incorrect status (%s)", s)
}

// convertToRatingState : 文字列を RatingState に変換
func convertToRatingState(s string) (gen.RatingState, error) {
	for _, rating := range gen.AllRatingState {
		if rating.String() == strings.ToUpper(s) {
			return rating, nil
		}
	}

	return "", fmt.Errorf("incorrect rating (%s)", s)
}

// ConvertToUpperFirstLetter : 頭文字を大文字に変換
func ConvertToUpperFirstLetter(s string) string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])

	return string(r)
}

// StripWhiteSpace : 空白文字を半角スペースに置換して前後をカット
func StripWhiteSpace(s string) string {
	// 全ての空白文字を半角スペースに置換
	r := regexp.MustCompile(`\s`)
	text := r.ReplaceAllString(s, " ")

	return strings.TrimSpace(text)
}
