package properties

import "github.com/pchchv/bpdf/consts/fontstyle"

// Font represents properties from a text.
type Font struct {
	Family string
	Style  fontstyle.Fontstyle
	Size   float64
	Color  *Color
}

// AppendMap appends the font fields to a map.
func (f *Font) AppendMap(m map[string]interface{}) map[string]interface{} {
	if f.Family != "" {
		m["prop_font_family"] = f.Family
	}

	if f.Style != "" {
		m["prop_font_style"] = f.Style
	}

	if f.Size != 0 {
		m["prop_font_size"] = f.Size
	}

	if f.Color != nil {
		m["prop_font_color"] = f.Color.ToString()
	}

	return m
}
