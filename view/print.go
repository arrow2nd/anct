package view

import (
	"fmt"
	"io"
	"text/template"

	"github.com/arrow2nd/anct/gen"
)

// PrintCanceled : キャンセル表示を出力
func PrintCanceled(w io.Writer) {
	fmt.Fprintln(w, "Canceled")
}

// PrintLogo : ロゴを出力
func PrintLogo(w io.Writer) {
	logo := `
   ________  ________  ________  ________ 
  /        \/    /   \/        \/        \
 /         /         /         /        _/
/         /         /       --//       /  
\___/____/\__/_____/\________/ \______/

`

	fmt.Fprint(w, logo)
}

// PrintAuthURL : 認証URLを出力
func PrintAuthURL(w io.Writer, u string) {
	temp := "📺 Please access the following URL and enter the code displayed after authentication.\n\n%s\n"
	fmt.Fprintf(w, temp, u)
}

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
   WEB:    {{url}}

DATA
----
   ANNICT ID: {{.AnnictID}}
   WORK ID:   {{.ID}}
   WATCHERS:  {{.WatchersCount}}
   STATUS:    {{.ViewerStatusState}}

{{if not .NoEpisodes -}}
EPISODES
--------
   {{if .Episodes.Nodes -}}
   {{range $i, $ep := .Episodes.Nodes -}}
   {{.NumberText}}  {{.Title}}
   {{end -}}
   {{else -}}
   None yet.
   {{end}}
{{- end}}
`

	t, err := template.New("work_info").Funcs(funcMap).Parse(temp)
	if err != nil {
		return err
	}

	// プロンプトの下に空行を挿入
	fmt.Fprintln(w)

	if err := printWorkImage(w, info.Image); err != nil {
		return err
	}

	return t.Execute(w, info)
}

// printWorkImage : 作品画像を出力
func printWorkImage(w io.Writer, workImage *gen.WorkInfoFragment_Image) error {
	if workImage == nil {
		return nil
	}

	imageURL := *workImage.RecommendedImageURL
	if imageURL == "" {
		imageURL = *workImage.FacebookOgImageURL
	}

	if imageURL == "" {
		return nil
	}

	if err := printImage(w, imageURL); err != nil {
		return err
	}

	return nil
}
