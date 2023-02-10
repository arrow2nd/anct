package view

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/arrow2nd/anct/gen"
	"github.com/ktr0731/go-fuzzyfinder"
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
func SelectWork(items []*gen.WorkFragment) (string, error) {
	idx, err := fuzzyfinder.Find(
		items,
		func(i int) string {
			return items[i].Title
		},
		fuzzyfinder.WithPreviewWindow(func(i, width, height int) string {
			if i == -1 {
				return ""
			}

			w := items[i]

			state := ""
			if s := w.ViewerStatusState; s != nil {
				state = fmt.Sprintf("  [ %s ]", *s)
			}

			media := "unknown"
			if w.Media.IsValid() {
				media = w.Media.String()
			}

			season := "unknown"
			if year := w.SeasonYear; w.SeasonName.IsValid() && year != nil {
				season = fmt.Sprintf("%s %d", w.SeasonName.String(), *year)
			}

			return fmt.Sprintf("%s%s\n\nID:     %s\nMEDIA:  %s\nSEASON: %s", w.Title, state, w.ID, media, season)
		}),
		fuzzyfinder.WithHeader("Please choose a work"),
	)

	if err != nil {
		return "", err
	}

	return items[idx].ID, nil
}
