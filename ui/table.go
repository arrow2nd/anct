package ui

import (
	"fmt"
	"os"

	"github.com/arrow2nd/anct/gen"
	"github.com/olekukonko/tablewriter"
)

func printTable(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(header)
	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetAutoMergeCells(true)
	table.AppendBulk(data)

	fmt.Println()
	table.Render()
	fmt.Println()
}

// PrintWorksTable : 作品テーブルを表示
func PrintWorksTable(works []*gen.WorkFragment) {
	if len(works) == 0 {
		fmt.Println("no matches found for your search")
		return
	}

	data := [][]string{}
	for _, work := range works {
		data = append(data, []string{
			fmt.Sprint(work.AnnictID),
			work.Title,
			work.Media.String(),
			fmt.Sprintf("%d %s", *work.SeasonYear, work.SeasonName.String()),
		})
	}

	printTable([]string{"WORK ID", "TITLE", "MEDIA", "SEASON"}, data)
}
