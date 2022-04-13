package result

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type tableViewResultRenderer struct {
}

func NewTableViewResultRenderer() *tableViewResultRenderer {
	return &tableViewResultRenderer{}
}

func (tvrr *tableViewResultRenderer) Render(results [][]string) {

	fmt.Println("")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"type", "total requests", "rps", "min", "max", "avg"})
	table.SetRowLine(true)

	for _, resultRow := range results {
		table.Append(resultRow)
	}

	table.Render()
}
