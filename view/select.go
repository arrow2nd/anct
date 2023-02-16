package view

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/arrow2nd/anct/gen"
)

// SelectStatus : 視聴ステータスを選択
func SelectStatus(allowNoState bool) (string, error) {
	opts := []string{}
	for _, status := range gen.AllStatusState {
		if !allowNoState && status == gen.StatusStateNoState {
			continue
		}
		opts = append(opts, string(status))
	}

	prompt := &survey.Select{
		Message: "Choose a status",
		Options: opts,
	}

	result := ""
	err := survey.AskOne(prompt, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

// SelectWork : 作品を選択
func SelectWork(works []*gen.WorkFragment) (int64, string, error) {
	opts := []string{}
	for _, work := range works {
		opts = append(opts, work.Title)
	}

	selectedTitle := ""
	prompt := &survey.Select{
		Message: "Choose a work",
		Options: opts,
		Description: func(_ string, index int) string {
			if index < 0 || index > len(works) {
				return ""
			}

			if s := works[index].ViewerStatusState; s != nil && *s != gen.StatusStateNoState {
				return s.String()
			}

			return ""
		},
	}

	if err := survey.AskOne(prompt, &selectedTitle); err != nil {
		return 0, "", err
	}

	for _, work := range works {
		if work.Title == selectedTitle {
			return work.AnnictID, work.ID, nil
		}
	}

	return 0, "", fmt.Errorf("Failed to retrieve the selected work ID (title: %s)", selectedTitle)
}
