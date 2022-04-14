package info

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/severuykhin/go-stress/interfaces"
	"github.com/severuykhin/go-stress/structs"
)

type tableViewInfoRenderer struct {
}

func NewTableViewInfoRenderer() *tableViewInfoRenderer {
	return &tableViewInfoRenderer{}
}

func (tvrr *tableViewInfoRenderer) Render(optFunc ...interfaces.InfoOptFunc) {
	table := tablewriter.NewWriter(os.Stdout)
	options := structs.InfoOptions{}

	for _, f := range optFunc {
		f(&options)
	}

	if options.Header != nil && len(options.Header) > 0 {
		table.SetHeader(options.Header)
	}

	table.SetRowLine(true)

	if options.Data != nil && len(options.Data) > 0 {
		for _, resultRow := range options.Data {
			table.Append(resultRow)
		}
	}

	table.Render()
}
