package properties

import (
	"github.com/pchchv/bpdf/consts/fontstyle"
	"github.com/pchchv/bpdf/consts/linestyle"
)

// Signature represents properties from a signature.
type Signature struct {
	FontFamily    string              // consts.Arial, helvetica and etc
	FontStyle     fontstyle.Fontstyle // consts.Normal, bold and etc
	FontSize      float64
	FontColor     *Color
	LineColor     *Color
	LineStyle     linestyle.LineStyle // solid or dashed
	LineThickness float64
	SafePadding   float64
}
