package properties

import "github.com/pchchv/bpdf/consts/fontstyle"

// Font represents properties from a text.
type Font struct {
	Family string
	Style  fontstyle.Fontstyle
	Size   float64
	Color  *Color
}
