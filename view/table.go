package view

import (
	"fmt"
	"io"

	"github.com/arrow2nd/anct/gen"
	"github.com/olekukonko/tablewriter"
)

func printTable(w io.Writer, header []string, data [][]string) {
	table := tablewriter.NewWriter(w)

	table.SetHeader(header)
	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetAutoWrapText(false)
	table.AppendBulk(data)

	fmt.Fprintln(w)
	table.Render()
	fmt.Fprintln(w)
}

// PrintWorksTable : 作品テーブルを出力
func PrintWorksTable(w io.Writer, q string, works []*gen.WorkFragment) {
	fmt.Fprintf(w, "\nSearch results for '%s' works\n", q)

	if len(works) == 0 {
		fmt.Fprintln(w, "No matches found for your search")
		return
	}

	data := [][]string{}
	for _, work := range works {
		media := "UNKNOWN"
		if work.Media.IsValid() {
			media = work.Media.String()
		}

		season := "UNKNOWN"
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

	printTable(w, []string{"WORK ID", "TITLE", "MEDIA", "SEASON"}, data)
}
