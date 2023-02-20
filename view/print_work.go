package view

import (
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/arrow2nd/anct/gen"
)

// PrintWorkInfo : 作品の詳細を出力
func PrintWorkInfo(w io.Writer, info *gen.WorkInfoFragment) error {
	funcMap := template.FuncMap{
		"season": func() string {
			s := ""

			if info.SeasonYear != nil {
				s += fmt.Sprintf("%d ", *info.SeasonYear)
			}
			if info.SeasonName != nil && info.SeasonName.IsValid() {
				s += info.SeasonName.String()
			}

			if s == "" {
				return "unknown"
			}

			return s
		},
		"url": func() string {
			if info.OfficialSiteURL == nil || *info.OfficialSiteURL == "" {
				return "unknown"
			}

			return *info.OfficialSiteURL
		},
	}

	temp := `DETAIL
------
   TITLE:  {{.Title}}
   MEDIA:  {{.Media}}
   SEASON: {{season}}
   URL:    {{url}}

DATA
----
   WORK ID:    {{.ID}}
   ANNICT ID:  {{.AnnictID}}
   ANNICT URL: https://annict.com/works/{{.AnnictID}}
   WATCHERS:   {{.WatchersCount}}
   STATUS:     {{.ViewerStatusState}}

EPISODES
--------
   {{if and (not .NoEpisodes) .Episodes.Nodes -}}
   {{range $i, $ep := .Episodes.Nodes -}}
   {{.NumberText}}  {{.Title}}
   {{end -}}
   {{else -}}
   None yet.
   {{end}}
`

	t, err := template.New("work_info").Funcs(funcMap).Parse(temp)
	if err != nil {
		return err
	}

	// プロンプトの下に空行を挿入
	fmt.Fprintln(w)

	if err := printWorkImage(w, info.Image); err != nil {
		// NOTE: 画像が表示できなくても処理は続けたいのでエラーを返さない
		fmt.Fprintf(os.Stderr, "failed to display image: %s", err.Error())
	}

	return t.Execute(w, info)
}

// printWorkImage : 作品画像を出力
func printWorkImage(w io.Writer, workImage *gen.WorkInfoFragment_Image) error {
	const imageResizeWidth = 400

	if workImage == nil {
		return nil
	}

	// 表示する画像のURLを決定
	imageURL := *workImage.RecommendedImageURL
	if imageURL == "" {
		imageURL = *workImage.FacebookOgImageURL
	}

	// 画像が無い場合は表示をスキップ
	if imageURL == "" {
		return nil
	}

	img, err := fetchImage(imageURL, imageResizeWidth)
	if err != nil {
		return fmt.Errorf("failed fetch image: %w", err)
	}

	if err := printImage(w, img); err != nil {
		return err
	}

	// コピーライト
	if workImage.Copyright != nil && *workImage.Copyright != "" {
		fmt.Fprintf(w, "(c) %s\n\n", *workImage.Copyright)
	}

	return nil
}
