package list

import (
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/srepio/sdk/types"
)

func NewTable(md *types.Metadata) table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Name", "Difficulty", "Version", "Tags")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, s := range *md {
		tbl.AddRow(s.Name, s.Difficulty, s.Version, strings.Join(s.Tags, ", "))
	}

	return tbl
}
