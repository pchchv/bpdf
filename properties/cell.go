package properties

import (
	"github.com/pchchv/bpdf/consts/border"
	"github.com/pchchv/bpdf/consts/linestyle"
)

// Cell is a representation of a cell in a grid system.
// It can be applied to Col or Row.
type Cell struct {
	// BackgroundColor defines which color will be applied to a cell.
	BackgroundColor *Color // Default: nil
	// BorderColor defines which color will be applied to a border cell
	BorderColor *Color // Default: nil
	// BorderType defines which kind of border will be applied to a cell.
	BorderType border.Border // Default: border.None
	// BorderThickness defines the border thickness applied to a cell.
	BorderThickness float64 // Default: 0.2
	// LineStyle defines which line style will be applied to a cell.
	LineStyle linestyle.LineStyle // Default: Solid
}

// ToMap adds the Cell fields to the map.
func (c *Cell) ToMap() map[string]interface{} {
	if c == nil {
		return nil
	}

	m := make(map[string]interface{})
	if c.BorderType != "" {
		m["prop_border_type"] = c.BorderType
	}

	if c.BorderThickness != 0 {
		m["prop_border_thickness"] = c.BorderThickness
	}

	if c.LineStyle != "" {
		m["prop_border_line_style"] = c.LineStyle
	}

	if c.BackgroundColor != nil {
		m["prop_background_color"] = c.BackgroundColor.ToString()
	}

	if c.BorderColor != nil {
		m["prop_border_color"] = c.BorderColor.ToString()
	}

	return m
}
