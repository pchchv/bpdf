// The properties package contains all props used to customize components.
package properties

import (
	"github.com/pchchv/bpdf/consts/align"
	"github.com/pchchv/bpdf/consts/breakline"
	"github.com/pchchv/bpdf/consts/fontstyle"
)

// Text represents properties from a Text inside a cell.
type Text struct {
	// Top is the amount of space between the upper cell limit and the text.
	Top float64
	// Bottom is the amount of space between the lower cell limit and the text.
	Bottom float64
	// Left is the minimal amount of space between the left cell boundary and the text.
	Left float64
	// Right is the minimal amount of space between the right cell boundary and the text.
	Right float64
	// Family of the text, ex: consts.Arial, helvetica and etc.
	Family string
	// Style of the text, ex: consts.Normal, bold and etc.
	Style fontstyle.Fontstyle
	// Size of the text.
	Size float64
	// Align of the text.
	Align align.Type
	// BreakLineStrategy define the break line strategy.
	BreakLineStrategy breakline.Strategy
	// VerticalPadding define an additional space between linet.
	VerticalPadding float64
	// Color define the font style color.
	Color *Color
	// Hyperlink define a link to be opened when the text is clicked.
	Hyperlink *string
}

// ToMap converts a Text to a map.
func (t *Text) ToMap() (m map[string]interface{}) {
	if t.Top != 0 {
		m["prop_top"] = t.Top
	}

	if t.Bottom != 0 {
		m["prop_bottom"] = t.Bottom
	}

	if t.Left != 0 {
		m["prop_left"] = t.Left
	}

	if t.Right != 0 {
		m["prop_right"] = t.Right
	}

	if t.Family != "" {
		m["prop_font_family"] = t.Family
	}

	if t.Style != "" {
		m["prop_font_style"] = t.Style
	}

	if t.Size != 0 {
		m["prop_font_size"] = t.Size
	}

	if t.Align != "" {
		m["prop_align"] = t.Align
	}

	if t.BreakLineStrategy != "" {
		m["prop_breakline_strategy"] = t.BreakLineStrategy
	}

	if t.VerticalPadding != 0 {
		m["prop_vertical_padding"] = t.VerticalPadding
	}

	if t.Color != nil {
		m["prop_color"] = t.Color.ToString()
	}

	if t.Hyperlink != nil {
		m["prop_hyperlink"] = *t.Hyperlink
	}

	return
}
