package gen

import (
	"bytes"
	"fmt"

	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
	"github.com/unidoc/unipdf/v3/model/optimize"
)

func renderText(c *creator.Creator, text string, textStyle *Text, fontStyle *Font) creator.VectorDrawable {
	desc := c.NewStyledParagraph()
	font, _ := model.NewStandard14Font(model.HelveticaName)
	t := desc.SetText(text)

	t.Style.Font = font
	t.Style.FontSize = 12

	if fontStyle != nil {
		if fontStyle.Color != "" {
			t.Style.Color = creator.ColorRGBFromHex(fontStyle.Color)
		}
	}

	if textStyle != nil {
		if textStyle.Align == "center" {
			desc.SetTextAlignment(creator.TextAlignmentCenter)
		}
	}

	return desc
}

func renderCell(c *creator.Creator, table *creator.Table, node *ElementNode) {
	fmt.Printf("Render cell\n")

	colspan := 1
	if node.Colspan > 0 {
		colspan = node.Colspan
	}
	cell := table.MultiColCell(colspan)

	bg := node.Background
	if bg != nil && bg.Color != "" {
		cell.SetBackgroundColor(creator.ColorRGBFromHex(bg.Color))
	}

	cell.SetContent(renderText(c, node.Data, node.Text, node.Font))
}

func renderRow(c *creator.Creator, table *creator.Table, node *ElementNode) {
	fmt.Printf("Render row\n")

	for _, cell := range node.Children {
		renderCell(c, table, cell)
	}
}

func renderTable(c *creator.Creator, node *ElementNode) creator.Drawable {
	tbody := node.Children[0]

	if len(tbody.Children) <= 0 {
		return nil
	}

	maxCols := 0
	firstRow := tbody.Children[0]
	for _, col := range firstRow.Children {
		if col.Colspan > 0 {
			maxCols += col.Colspan
		} else {
			maxCols += 1
		}
	}

	fmt.Printf("maxCols: %v\n", maxCols)
	table := c.NewTable(maxCols)

	for _, row := range tbody.Children {
		renderRow(c, table, row)
	}

	return table
}

var renders = map[string]func(*creator.Creator, *ElementNode) creator.Drawable{
	"table": renderTable,
}

func render(doc *ElementNode) []byte {
	c := creator.New()

	for _, child := range doc.Children {
		ele := renderElement(c, child)
		if ele != nil {
			if err := c.Draw(ele); err != nil {
				fmt.Printf("Failed to draw element: %v\n", err)
			}
		}
	}

	optimizer := optimize.New(optimize.Options{
		CombineDuplicateStreams:         true,
		CombineDuplicateDirectObjects:   true,
		ImageUpperPPI:                   100.0,
		ImageQuality:                    90,
		CombineIdenticalIndirectObjects: true,
		UseObjectStreams:                true,
		CompressStreams:                 true,
	})

	c.SetOptimizer(optimizer)

	var res bytes.Buffer
	if err := c.Write(&res); err != nil {
		fmt.Printf("Failed to write bytes: %v\n", err)
	}

	return res.Bytes()
}

func renderElement(c *creator.Creator, node *ElementNode) creator.Drawable {
	if process, ok := renders[node.Type]; ok {
		return process(c, node)
	}
	return nil
}
