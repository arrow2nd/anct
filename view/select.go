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

	// 選択項目を作成
	for _, status := range gen.AllStatusState {
		if !allowNoState && status == gen.StatusStateNoState {
			continue
		}
		opts = append(opts, status.String())
	}

	prompt := &survey.Select{
		Message: "Status",
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
func SelectRating(m string) (string, error) {
	opts := []string{}
	for _, rating := range gen.AllRatingState {
		opts = append(opts, rating.String())
	}

	prompt := &survey.Select{
		Message: m,
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
		Message: "Work",
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

	return 0, "", fmt.Errorf("failed to retrieve the selected work ID (title: %s)", selectedTitle)
}

// SelectEpisodes : エピソードを選択
func SelectEpisodes(work *gen.WorkEpisodesFragment) ([]string, error) {
	// エピソードが無い作品
	if work.NoEpisodes || work.Episodes == nil || len(work.Episodes.Nodes) == 0 {
		return nil, errors.New("no selectable episodes")
	}

	opts := []string{}
	for _, ep := range work.Episodes.Nodes {
		opts = append(opts, createEpisodeOpt(ep, true))
	}

	prompt := &survey.MultiSelect{
		Message: "Episode",
		Options: opts,
	}

	selectedOpts := []string{}
	if err := survey.AskOne(prompt, &selectedOpts); err != nil {
		return nil, err
	}

	// 選択した項目をエピソードIDの配列に変換
	episodeIDs := []string{}
	for _, opt := range selectedOpts {
		for _, e := range work.Episodes.Nodes {
			if opt == createEpisodeOpt(e, true) {
				episodeIDs = append(episodeIDs, e.ID)
				break
			}
		}
	}

	// 変換前と件数が一致しない場合エラー
	if before, after := len(selectedOpts), len(episodeIDs); before != after {
		return nil, fmt.Errorf("failed to obtain episode ID (%d of %d failed)", before-after, before)
	}

	return episodeIDs, nil
}

// SelectUnwatchEpisode : 未視聴のエピソードを選択
func SelectUnwatchEpisode(entries []*gen.UnwatchLibraryEntryFragment) (string, error) {
	opts := []string{}
	for _, entry := range entries {
		opts = append(opts, createEpisodeOpt(entry.NextEpisode, false))
	}

	prompt := &survey.Select{
		Message: "Episode",
		Options: opts,
		Description: func(_ string, i int) string {
			if i < 0 || i > len(entries) {
				return ""
			}

			return entries[i].Work.Title
		},
	}

	selectedOpt := ""
	if err := survey.AskOne(prompt, &selectedOpt); err != nil {
		return "", err
	}

	// エピソードIDを返す
	for _, e := range entries {
		if selectedOpt == createEpisodeOpt(e.NextEpisode, true) {
			return e.NextEpisode.ID, nil
		}
	}

	return "", fmt.Errorf("failed to obtain episode ID (selected: %s)", selectedOpt)
}

// createEpisodeOpt : エピソードの選択項目文字列を作成
func createEpisodeOpt(e *gen.EpisodeFragment, showRecorded bool) string {
	// 話数
	num := "???"
	if e.NumberText != nil && *e.NumberText != "" {
		num = *e.NumberText
	}

	// タイトル
	title := fmt.Sprintf("??? (ID: %s)", e.ID)
	if e.Title != nil && *e.Title != "" {
		title = *e.Title
	}

	// 記録済なら件数をタイトル末尾に追加
	if showRecorded && e.ViewerRecordsCount != 0 {
		title += fmt.Sprintf(" - Recorded (%d)", e.ViewerRecordsCount)
	}

	return fmt.Sprintf("%s %s", num, title)
}
