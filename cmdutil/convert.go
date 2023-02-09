package cmdutil

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/arrow2nd/anct/gen"
)

// StringToStatusState : 文字列を視聴ステータスに変換
func StringToStatusState(s string, allowNoState bool) (gen.StatusState, error) {
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

// StringToWorkID : 文字列を Work ID に 変換
func StringToWorkID(s string) (string, error) {
	id := strings.TrimSpace(s)

	r := regexp.MustCompile(`^\d+$`)
	if !r.MatchString(id) {
		return "", errors.New("work id must be numeric")
	}

	return id, nil
}
