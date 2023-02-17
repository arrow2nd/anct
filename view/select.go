package view

import (
	"errors"
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
		opts = append(opts, status.String())
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

// SelectRating : 評価を選択
func SelectRating() (string, error) {
	opts := []string{}
	for _, rating := range gen.AllRatingState {
		opts = append(opts, rating.String())
	}

	prompt := &survey.Select{
		Message: "Choose a rating",
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

	selectedTitle := ""
	if err := survey.AskOne(prompt, &selectedTitle); err != nil {
		return 0, "", err
	}

	// 選択した作品のIDを返す
	for _, work := range works {
		if work.Title == selectedTitle {
			return work.AnnictID, work.ID, nil
		}
	}

	return 0, "", fmt.Errorf("Failed to retrieve the selected work ID (title: %s)", selectedTitle)
}

// SelectEpisodes : エピソードを選択
func SelectEpisodes(work *gen.WorkEpisodesFragment) ([]string, error) {
	// エピソードが無い作品
	if work.NoEpisodes || work.Episodes == nil || len(work.Episodes.Nodes) == 0 {
		return nil, errors.New("no selectable episodes")
	}

	opts := []string{}
	for _, ep := range work.Episodes.Nodes {
		opts = append(opts, createEpisodeOpt(ep))
	}

	prompt := &survey.MultiSelect{
		Message: "Choose a episode",
		Options: opts,
	}

	selectedOpts := []string{}
	if err := survey.AskOne(prompt, &selectedOpts); err != nil {
		return nil, err
	}

	// 選択した項目をエピソードIDの配列に変換
	episodeIDs := []string{}
	for _, opt := range selectedOpts {
		for _, ep := range work.Episodes.Nodes {
			if opt == createEpisodeOpt(ep) {
				episodeIDs = append(episodeIDs, ep.ID)
				break
			}
		}
	}

	return episodeIDs, nil
}

// createEpisodeOpt : エピソードの選択項目文字列を作成
func createEpisodeOpt(e *gen.WorkEpisodesFragment_Episodes_Nodes) string {
	num := "???"
	if e.NumberText != nil && *e.NumberText != "" {
		num = *e.NumberText
	}

	title := fmt.Sprintf("??? (ID: %s)", e.ID)
	if e.Title != nil && *e.Title != "" {
		title = *e.Title
	}

	// 記録済なら件数を追加
	if e.ViewerRecordsCount != 0 {
		title += fmt.Sprintf(" - Recorded (%d)", e.ViewerRecordsCount)
	}

	return fmt.Sprintf("%s %s", num, title)
}
