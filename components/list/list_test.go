package list_test

import (
	"github.com/pchchv/bpdf/components/row"
	"github.com/pchchv/bpdf/components/text"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/internal/fixture"
)

type anyType struct {
	Key   string
	Value string
}

func (a anyType) GetHeader() core.Row {
	r := row.New(10).Add(
		text.NewCol(6, "Key"),
		text.NewCol(6, "Value"),
	)

	return r
}

func (a anyType) GetContent(i int) core.Row {
	r := row.New(10).Add(
		text.NewCol(6, a.Key),
		text.NewCol(6, a.Value),
	)

	if i%2 == 0 {
		cell := fixture.CellProp()
		r.WithStyle(&cell)
	}

	return r
}
