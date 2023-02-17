package cmdutil

import (
	"fmt"
	"strings"

	"github.com/arrow2nd/anct/gen"
)

// StringToStatusState : 文字列を StatusState に変換
func StringToStatusState(s string, allowNoState bool) (gen.StatusState, error) {
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

// StringToRatingState : 文字列を RatingState に変換
func StringToRatingState(s string) (gen.RatingState, error) {
	for _, rating := range gen.AllRatingState {
		if rating.String() == strings.ToUpper(s) {
			return rating, nil
		}
	}

	return "", fmt.Errorf("incorrect rating (%s)", s)
}
