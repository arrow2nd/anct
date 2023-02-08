package view

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
	"github.com/arrow2nd/anct/gen"
)

// SelectStatus : 視聴ステータスを選択
func SelectStatus() (gen.StatusState, error) {
	opts := []string{}
	for _, status := range gen.AllStatusState {
		opts = append(opts, string(status))
	}

	prompt := &survey.Select{
		Message: "Choose a status",
		Options: opts,
	}

	result := ""
	err := survey.AskOne(prompt, &result)
	if err != nil {
		return gen.StatusStateNoState, err
	}

	for _, status := range gen.AllStatusState {
		if string(status) == result {
			return status, nil
		}
	}

	return gen.StatusStateNoState, errors.New("failed to retrieve selection results")
}