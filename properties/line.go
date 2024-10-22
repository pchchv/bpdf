package properties

import (
	"github.com/pchchv/bpdf/consts/linestyle"
	"github.com/pchchv/bpdf/consts/orientation"
)

// Line represents properties from a Line inside a cell.
type Line struct {
	Color         *Color
	Style         linestyle.LineStyle // solid or dashed
	Thickness     float64
	Orientation   orientation.Orient // horizontal or vertical
	OffsetPercent float64            // line position, 0 - start of cell, 50 - middle, 100 - end
	SizePercent   float64
}
