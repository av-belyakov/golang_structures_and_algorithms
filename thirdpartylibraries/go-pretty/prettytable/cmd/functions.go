package prettytable

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func PrettyTableExample(footerNames []any, rows [][]any) {
	fmt.Println("func 'PrettyTableExample'")

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.Style{
		Name: "myNewStyle",
		Box: table.BoxStyle{
			BottomLeft:       "\\",
			BottomRight:      "/",
			BottomSeparator:  "v",
			Left:             "[",
			LeftSeparator:    "{",
			MiddleHorizontal: "-",
			MiddleSeparator:  "+",
			MiddleVertical:   "|",
			PaddingLeft:      "<",
			PaddingRight:     ">",
			Right:            "]",
			RightSeparator:   "}",
			TopLeft:          "(",
			TopRight:         ")",
			TopSeparator:     "^",
			UnfinishedRow:    " ~~~",
		},
		Color: table.ColorOptions{
			IndexColumn:  text.Colors{text.BgCyan, text.FgBlack},
			Footer:       text.Colors{text.BgCyan, text.FgBlack},
			Header:       text.Colors{text.BgHiGreen, text.FgBlack},
			Row:          text.Colors{text.BgHiWhite, text.FgBlack},
			RowAlternate: text.Colors{text.BgWhite, text.FgBlack},
		},
		Format: table.FormatOptions{
			Footer: text.FormatUpper,
			Header: text.FormatUpper,
			Row:    text.FormatDefault,
		},
		Options: table.Options{
			DrawBorder:      true,
			SeparateColumns: true,
			SeparateFooter:  true,
			SeparateHeader:  true,
			SeparateRows:    false,
		},
	})

	//t.AppendHeader(table.Row{headerNames})
	t.AppendHeader(table.Row{
		"id",
		"Name",
		"Date",
		"Address",
		"Cost",
	})
	t.AppendSeparator()

	for _, v := range rows {
		t.AppendRows([]table.Row{v})
	}

	t.AppendFooter(footerNames)
	t.Render()
}
