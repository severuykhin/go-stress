package info

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/severuykhin/go-stress/structs"
)

type tableViewInfoRenderer struct {
}

func NewTableViewInfoRenderer() *tableViewInfoRenderer {
	return &tableViewInfoRenderer{}
}

func (tvrr *tableViewInfoRenderer) Render(infoData structs.InfoData) {
	table := tablewriter.NewWriter(os.Stdout)

	if infoData.Header != nil && len(infoData.Header) > 0 {
		table.SetHeader(infoData.Header)
	}

	table.SetRowLine(true)

	if infoData.Data != nil && len(infoData.Data) > 0 {
		for _, resultRow := range infoData.Data {
			table.Append(resultRow)
		}
	}

	table.Render()
}
