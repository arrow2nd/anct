package view

import (
	"fmt"
	"io"

	"github.com/arrow2nd/anct/gen"
	"github.com/olekukonko/tablewriter"
)

func printTable(w io.Writer, q string, header []string, data [][]string) {
	table := tablewriter.NewWriter(w)

	table.SetHeader(header)
	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetAutoWrapText(false)
	table.AppendBulk(data)

	fmt.Fprintf(w, "\nSearch results for '%s'\n\n", q)
	table.Render()
	fmt.Fprintln(w)
}

// PrintWorksTable : 作品テーブルを出力
func PrintWorksTable(w io.Writer, q string, works []*gen.WorkFragment) {
	if len(works) == 0 {
		fmt.Fprintln(w, "No matches found for your search")
		return
	}

	data := [][]string{}
	for _, work := range works {
		media := "?"
		if work.Media.IsValid() {
			media = work.Media.String()
		}

		season := "?"
		if work.SeasonYear != nil && work.SeasonName.IsValid() {
			season = fmt.Sprintf("%d %s", *work.SeasonYear, work.SeasonName.String())
		}

		data = append(data, []string{
			fmt.Sprint(work.AnnictID),
			work.Title,
			media,
			season,
		})
	}

	printTable(w, q, []string{"WORK ID", "TITLE", "MEDIA", "SEASON"}, data)
}

func PrintCharactersTable(w io.Writer, q string, charachers []*gen.CharacterFragment) {
	if len(charachers) == 0 {
		fmt.Fprintln(w, "No matches found for your search")
		return
	}

	data := [][]string{}
	for _, chara := range charachers {
		data = append(data, []string{
			chara.Name,
			chara.Series.Name,
			fmt.Sprint(chara.Series.AnnictID),
		})
	}

	printTable(w, q, []string{"NAME", "SERIES", "WORK ID"}, data)
}
