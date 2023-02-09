package cmdutil

import (
	"errors"
	"fmt"
	"strconv"
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
func StringToWorkID(s string) (int64, error) {
	workID, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("work id must be numeric")
	}

	if workID <= 0 {
		return 0, errors.New("work id must be greater than or equal to 1")
	}

	return int64(workID), nil
}
